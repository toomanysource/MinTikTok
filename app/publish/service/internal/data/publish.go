package data

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/url"
	"os"
	"strconv"
	"sync"
	"time"

	"Atreus/app/publish/service/internal/biz"
	"Atreus/app/publish/service/internal/server"
	"Atreus/pkg/ffmpegX"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
)

// Video Database Model
type Video struct {
	Id            uint32 `gorm:"column:id;primary_key;auto_increment"`
	AuthorID      uint32 `gorm:"column:author_id;not null"`
	Title         string `gorm:"column:title;not null;size:255"`
	PlayUrl       string `gorm:"column:play_url;not null"`
	CoverUrl      string `gorm:"column:cover_url;not null"`
	FavoriteCount uint32 `gorm:"column:favorite_count;not null;default:0"`
	CommentCount  uint32 `gorm:"column:comment_count;not null;default:0"`
	CreatedAt     int64  `gorm:"column:created_at"`
}

type UserRepo interface {
	GetUserInfos(context.Context, uint32, []uint32) ([]*biz.User, error)
	UpdateVideoCount(context.Context, uint32, int32) error
}
type FavoriteRepo interface {
	IsFavorite(context.Context, uint32, []uint32) ([]bool, error)
}

type publishRepo struct {
	data         *Data
	favoriteRepo FavoriteRepo
	userRepo     UserRepo
	log          *log.Helper
}

func NewPublishRepo(
	data *Data, userConn server.UserConn, favoriteConn server.FavoriteConn, logger log.Logger,
) biz.PublishRepo {
	return &publishRepo{
		data:         data,
		favoriteRepo: NewFavoriteRepo(favoriteConn),
		userRepo:     NewUserRepo(userConn),
		log:          log.NewHelper(logger),
	}
}

// UploadVideo 上传视频
func (r *publishRepo) UploadVideo(ctx context.Context, fileBytes []byte, userId uint32, title string) error {
	return r.data.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Where("author_id = ? AND title = ?", userId, title).First(&Video{}).Error
		if err == nil {
			return fmt.Errorf("video already exists")
		}
		var wg sync.WaitGroup
		wg.Add(2)
		errChan := make(chan error)
		// 生成封面
		go func() {
			defer wg.Done()
			coverReader, err := r.GenerateCoverImage(fileBytes)
			if err != nil {
				errChan <- fmt.Errorf("generate cover image error: %w", err)
				return
			}
			data, err := io.ReadAll(coverReader)
			if err != nil {
				errChan <- fmt.Errorf("read cover image error: %w", err)
				return
			}
			coverBytes := bytes.NewReader(data)
			// 上传封面
			err = r.data.oss.UploadSizeFile(
				ctx, "oss", "images/"+title+".png", coverBytes, coverBytes.Size(), minio.PutObjectOptions{
					ContentType: "image/png",
				})
			if err != nil {
				errChan <- fmt.Errorf("upload cover image error: %w", err)
				return
			}
		}()

		// 上传视频
		go func() {
			defer wg.Done()
			reader := bytes.NewReader(fileBytes)
			err = r.data.oss.UploadSizeFile(
				ctx, "oss", "videos/"+title+".mp4", reader, reader.Size(), minio.PutObjectOptions{
					ContentType: "video/mp4",
				})
			if err != nil {
				errChan <- fmt.Errorf("upload video error: %w", err)
				return
			}
		}()
		wg.Wait()
		select {
		case err = <-errChan:
			return err
		default:
			// 获取视频和封面的url
			playUrl, coverUrl, err := r.GetRemoteVideoInfo(ctx, title)
			if err != nil {
				return fmt.Errorf("get remote video info error: %w", err)
			}
			v := &Video{
				AuthorID:      userId,
				Title:         title,
				PlayUrl:       playUrl,
				CoverUrl:      coverUrl,
				FavoriteCount: 0,
				CommentCount:  0,
				CreatedAt:     time.Now().UnixMilli(),
			}
			if err := tx.Create(v).Error; err != nil {
				return fmt.Errorf("create video error: %w", err)
			}
		}
		err = r.userRepo.UpdateVideoCount(ctx, userId, 1)
		if err != nil {
			return fmt.Errorf("update user video count error: %w", err)
		}
		return nil
	})
}

// GetRemoteVideoInfo 获取远程视频信息
func (r *publishRepo) GetRemoteVideoInfo(ctx context.Context, title string) (playURL, coverURL string, err error) {
	urls, err := r.data.oss.GetFileURL(
		ctx, "oss", "videos/"+title+".mp4", time.Hour*24*7)
	if err != nil {
		return "", "", fmt.Errorf("get video url error: %w", err)
	}
	playURL = urls.String()
	urls, err = r.data.oss.GetFileURL(
		ctx, "oss", "images/"+title+".png", time.Hour*24*7)
	if err != nil {
		return "", "", fmt.Errorf("get image url error: %w", err)
	}
	coverURL = urls.String()
	return
}

// GenerateCoverImage 生成封面
func (r *publishRepo) GenerateCoverImage(fileBytes []byte) (io.Reader, error) {
	// 创建临时文件
	tempFile, err := os.CreateTemp("", "tempFile-*")
	if err != nil {
		return nil, fmt.Errorf("create temp file error: %w", err)
	}
	defer os.Remove(tempFile.Name())
	if _, err = tempFile.Write(fileBytes); err != nil {
		return nil, fmt.Errorf("write temp file error: %w", err)
	}
	// 调用ffmpeg 生成封面
	return ffmpegX.ReadFrameAsImage(tempFile.Name(), 60)
}

func (r *publishRepo) FindVideoListByUserId(ctx context.Context, userId uint32) ([]*biz.Video, error) {
	var videoList []*Video
	result := r.data.db.WithContext(ctx).Where("author_id = ?", userId).Find(&videoList)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	err := r.UpdateUrl(ctx, videoList)
	if err != nil {
		return nil, fmt.Errorf("update url error: %w", err)
	}
	users, err := r.userRepo.GetUserInfos(ctx, 0, []uint32{userId})
	if err != nil {
		return nil, err
	}
	videoIds := make([]uint32, 0, len(videoList))
	for _, video := range videoList {
		videoIds = append(videoIds, video.Id)
	}
	isFavoriteList, err := r.favoriteRepo.IsFavorite(ctx, userId, videoIds)
	if err != nil {
		return nil, err
	}
	vl := make([]*biz.Video, 0, len(videoList))
	for i, video := range videoList {
		vl = append(vl, &biz.Video{
			ID:            video.Id,
			Author:        users[0],
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    isFavoriteList[i],
			Title:         video.Title,
		})
	}
	return vl, nil
}

func (r *publishRepo) FindVideoListByVideoIds(ctx context.Context, userId uint32, videoIds []uint32) ([]*biz.Video, error) {
	if len(videoIds) == 0 {
		return nil, nil
	}
	var videoList []*Video
	err := r.data.db.WithContext(ctx).Where("id IN ?", videoIds).Find(&videoList).Error
	if err != nil {
		return nil, err
	}
	err = r.UpdateUrl(ctx, videoList)
	if err != nil {
		return nil, fmt.Errorf("update url error: %w", err)
	}
	return r.GetVideoAuthor(ctx, userId, videoList)
}

func (r *publishRepo) FindVideoListByTime(
	ctx context.Context, latestTime string, userId uint32, number uint32,
) (int64, []*biz.Video, error) {
	var videoList []*Video
	times, err := strconv.Atoi(latestTime)
	if err != nil {
		return 0, nil, fmt.Errorf("strconv.Atoi error: %w", err)
	}
	err = r.data.db.WithContext(ctx).Where("created_at < ?", int64(times)).
		Order("created_at desc").Limit(int(number)).Find(&videoList).Error
	if err != nil {
		return 0, nil, fmt.Errorf("find video list error: %w", err)
	}
	if len(videoList) == 0 {
		return 0, nil, nil
	}
	err = r.UpdateUrl(ctx, videoList)
	if err != nil {
		return 0, nil, fmt.Errorf("update url error: %w", err)
	}
	nextTime := videoList[len(videoList)-1].CreatedAt
	vl, err := r.GetVideoAuthor(ctx, userId, videoList)
	if err != nil {
		return 0, nil, fmt.Errorf("get users error: %w", err)
	}

	// userId == 0 代表未登录
	if userId != 0 {
		videoIds := make([]uint32, 0, len(videoList))
		for _, video := range videoList {
			videoIds = append(videoIds, video.Id)
		}
		isFavoriteList, err := r.favoriteRepo.IsFavorite(ctx, userId, videoIds)
		if err != nil {
			return 0, nil, err
		}
		for i, video := range vl {
			video.IsFavorite = isFavoriteList[i]
		}
		return nextTime, vl, err
	}
	for _, video := range vl {
		video.IsFavorite = false
	}
	return nextTime, vl, err
}

func (r *publishRepo) GetVideoAuthor(ctx context.Context, userId uint32, videoList []*Video) ([]*biz.Video, error) {
	userIds := make([]uint32, 0, len(videoList))
	for _, video := range videoList {
		userIds = append(userIds, video.AuthorID)
	}
	userMap := make(map[uint32]*biz.User)
	users, err := r.userRepo.GetUserInfos(ctx, userId, userIds)
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		userMap[user.ID] = user
	}
	vl := make([]*biz.Video, 0, len(videoList))
	for i, video := range videoList {
		vl = append(vl, &biz.Video{
			ID:            video.Id,
			Author:        userMap[video.AuthorID],
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    false,
			Title:         video.Title,
		})
	}
	return vl, nil
}

func (r *publishRepo) UpdateFavoriteCount(ctx context.Context, videoId uint32, favoriteChange int32) error {
	var video Video
	err := r.data.db.WithContext(ctx).Where("id = ?", videoId).First(&video).Error
	if err != nil {
		return err
	}
	newCount := calculateValidUint32(video.FavoriteCount, favoriteChange)
	err = r.data.db.Model(&Video{}).Where("id = ?", videoId).Update("favorite_count", newCount).Error
	if err != nil {
		return err
	}
	return err
}

func (r *publishRepo) UpdateCommentCount(ctx context.Context, videoId uint32, commentChange int32) error {
	var video Video
	err := r.data.db.WithContext(ctx).Where("id = ?", videoId).First(&video).Error
	if err != nil {
		return err
	}
	newCount := calculateValidUint32(video.FavoriteCount, commentChange)
	err = r.data.db.Model(&Video{}).Where("id = ?", videoId).
		Update("comment_count", newCount).Error
	if err != nil {
		return err
	}
	return err
}

// CheckUrl 检查url是否过期
func (r *publishRepo) CheckUrl(accessUrl string) (bool, error) {
	parseUrl, err := url.Parse(accessUrl)
	if err != nil {
		return false, fmt.Errorf("parse url error: %w", err)
	}
	dateStr := parseUrl.Query().Get("X-Amz-Date")
	dateInt, err := time.Parse("20060102T150405Z", dateStr)
	if err != nil {
		return false, fmt.Errorf("parse date error: %w", err)
	}
	// 7天后过期,提前一个小时生成新的url
	now := time.Now().Add(time.Hour).UTC()
	if now.After(dateInt.Add(time.Hour * 24 * 7)) {
		return false, nil
	}
	return true, nil
}

// UpdateUrl 更新url
func (r *publishRepo) UpdateUrl(ctx context.Context, videoList []*Video) error {
	for _, video := range videoList {
		ok, err := r.CheckUrl(video.PlayUrl)
		if err != nil {
			return err
		}
		if !ok {
			video.PlayUrl, video.CoverUrl, err = r.GetRemoteVideoInfo(ctx, video.Title)
			if err != nil {
				return err
			}
			err = r.UpdateDatabaseUrl(ctx, video.Id, video.PlayUrl, video.CoverUrl)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// UpdateDatabaseUrl 更新数据库url
func (r *publishRepo) UpdateDatabaseUrl(ctx context.Context, videoId uint32, playUrl, coverUrl string) error {
	err := r.data.db.WithContext(ctx).Where("id = ?", videoId).
		Updates(&Video{PlayUrl: playUrl, CoverUrl: coverUrl}).Error
	if err != nil {
		return fmt.Errorf("update database url error: %w", err)
	}
	return nil
}

func calculateValidUint32(src uint32, mod int32) uint32 {
	if mod < 0 {
		mod = -mod
		if src < uint32(mod) {
			return 0
		}
		return src - uint32(mod)
	}
	return src + uint32(mod)
}
