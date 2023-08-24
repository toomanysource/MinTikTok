package data

import (
	"context"
	"strconv"
	"time"

	"Atreus/app/feed/service/internal/biz"
	"Atreus/app/feed/service/internal/server"

	"github.com/go-kratos/kratos/v2/log"
)

var VideoCount uint32 = 30

type PublishRepo interface {
	GetVideoList(context.Context, string, uint32, uint32) (int64, []*biz.Video, error)
}

type feedRepo struct {
	publishRepo biz.PublishRepo
	log         *log.Helper
}

func NewFeedRepo(publishConn server.PublishConn, logger log.Logger) biz.FeedRepo {
	return &feedRepo{
		publishRepo: NewPublishRepo(publishConn),
		log:         log.NewHelper(log.With(logger, "model", "data/feed")),
	}
}

// GetFeedList 获取feed列表
func (r *feedRepo) GetFeedList(ctx context.Context, latestTime string, userId uint32) (nextTime int64, vl []*biz.Video, err error) {
	if latestTime == "0" {
		latestTime = strconv.FormatInt(time.Now().UnixMilli(), 10)
	}
	switch userId {
	case 0:
		return r.publishRepo.GetVideoList(ctx, latestTime, 0, VideoCount)
	default:
		return r.publishRepo.GetVideoList(ctx, latestTime, userId, VideoCount)
	}
}
