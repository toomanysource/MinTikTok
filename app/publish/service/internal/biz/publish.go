package biz

import (
	"Atreus/app/publish/service/internal/conf"
	"Atreus/pkg/common"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

// Video is a video model
type Video struct {
	ID            uint32
	Author        *User
	PlayUrl       string
	CoverUrl      string
	FavoriteCount uint32
	CommentCount  uint32
	IsFavorite    bool
	Title         string
}

// User is a user model.
type User struct {
	ID              uint32
	Name            string
	FollowCount     uint32
	FollowerCount   uint32
	IsFollow        bool
	Avatar          string
	BackgroundImage string
	Signature       string
	TotalFavorited  uint32
	WorkCount       uint32
	FavoriteCount   uint32
}

// PublishRepo is a publishing repo.
type PublishRepo interface {
	FindVideoListByUserId(context.Context, uint32) ([]*Video, error)
	UploadVideo(context.Context, []byte, uint32, string) error
	FindVideoListByTime(context.Context, string, uint32, uint32) (int64, []*Video, error)
	FindVideoListByIDs(context.Context, uint32, []uint32) ([]*Video, error)
	UpdateFavoriteCount(context.Context, uint32, int32) error
	UpdateCommentCount(context.Context, uint32, int32) error
}

// PublishUsecase is a publishing usecase.
type PublishUsecase struct {
	repo   PublishRepo
	config *conf.JWT
	log    *log.Helper
}

// NewPublishUsecase new a publishing usecase.
func NewPublishUsecase(repo PublishRepo, JWTConf *conf.JWT, logger log.Logger) *PublishUsecase {
	return &PublishUsecase{repo: repo, config: JWTConf, log: log.NewHelper(logger)}
}

func (u *PublishUsecase) GetPublishList(
	ctx context.Context, tokenString string, userId uint32) ([]*Video, error) {
	if tokenString != "" {
		token, err := common.ParseToken(u.config.Http.TokenKey, tokenString)
		if err != nil {
			return nil, err
		}
		_, err = common.GetTokenData(token)
		if err != nil {
			return nil, err
		}
	}
	return u.repo.FindVideoListByUserId(ctx, userId)
}

func (u *PublishUsecase) PublishAction(
	ctx context.Context, fileBytes []byte, title, tokenString string) error {
	token, err := common.ParseToken(u.config.Http.TokenKey, tokenString)
	if err != nil {
		return err
	}
	data, err := common.GetTokenData(token)
	if err != nil {
		return err
	}
	userId := uint32(data["user_id"].(float64))
	return u.repo.UploadVideo(ctx, fileBytes, userId, title)
}

func (u *PublishUsecase) GetVideoListByVideoIds(ctx context.Context, userId uint32, ids []uint32) ([]*Video, error) {
	videoList, err := u.repo.FindVideoListByIDs(ctx, userId, ids)
	return videoList, err
}

func (u *PublishUsecase) GetVideoList(
	ctx context.Context, latestTime string, userId uint32, number uint32) (int64, []*Video, error) {
	return u.repo.FindVideoListByTime(ctx, latestTime, userId, number)
}

func (u *PublishUsecase) UpdateFavorite(ctx context.Context, videoId uint32, favoriteChange int32) error {
	return u.repo.UpdateFavoriteCount(ctx, videoId, favoriteChange)
}

func (u *PublishUsecase) UpdateComment(ctx context.Context, videoId uint32, commentChange int32) error {
	return u.repo.UpdateCommentCount(ctx, videoId, commentChange)
}
