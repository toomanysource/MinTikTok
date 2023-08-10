package biz

import (
	"Atreus/app/relation/service/internal/conf"
	"Atreus/pkg/common"
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Id              uint32 // 用户id
	Name            string // 用户名称
	FollowCount     uint32 // 关注总数
	FollowerCount   uint32 // 粉丝总数
	IsFollow        bool   // true-已关注，false-未关注
	Avatar          string //用户头像
	BackgroundImage string //用户个人页顶部大图
	Signature       string //个人简介
	TotalFavorite   uint32 //获赞数量
	WorkCount       uint32 //作品数量
	FavoriteCount   uint32 //点赞数量
}

type RelationRepo interface {
	GetFollowList(context.Context, uint32) ([]*User, error)
	GetFollowerList(context.Context, uint32) ([]*User, error)
	Follow(context.Context, uint32, uint32) error
	UnFollow(context.Context, uint32, uint32) error
	GetFollowCount(context.Context, uint32) (int64, error)
	GetFollowerCount(context.Context, uint32) (int64, error)
}

type RelationUsecase struct {
	repo   RelationRepo
	config *conf.JWT
	log    *log.Helper
}

func NewRelationUsecase(repo RelationRepo, logger log.Logger) *RelationUsecase {
	return &RelationUsecase{repo: repo, log: log.NewHelper(logger)}
}

// GetFollowList 获取关注列表
func (uc *RelationUsecase) GetFollowList(ctx context.Context, userId uint32, tokenString string) ([]*User, error) {
	token, err := common.ParseToken(uc.config.Http.TokenKey, tokenString)
	if err != nil {
		return nil, err
	}
	_, err = common.GetTokenData(token)
	if err != nil {
		return nil, err
	}
	return uc.repo.GetFollowList(ctx, userId)
}

// GetFollowerList 获取粉丝列表
func (uc *RelationUsecase) GetFollowerList(ctx context.Context, userId uint32, tokenString string) ([]*User, error) {
	token, err := common.ParseToken(uc.config.Http.TokenKey, tokenString)
	if err != nil {
		return nil, err
	}
	_, err = common.GetTokenData(token)
	if err != nil {
		return nil, err
	}
	return uc.repo.GetFollowerList(ctx, userId)
}

// Action 关注和取消关注
func (uc *RelationUsecase) Action(ctx context.Context, tokenString string, toUserId uint32, actionType uint32) error {
	token, err := common.ParseToken(uc.config.Http.TokenKey, tokenString)
	if err != nil {
		return err
	}
	data, err := common.GetTokenData(token)
	if err != nil {
		return err
	}
	userId := uint32(data["user_id"].(float64))
	switch actionType {
	//1为关注
	case 1:
		err := uc.repo.Follow(ctx, userId, toUserId)
		if err != nil {
			return errors.New("something wrong")
		}
	//2为取消关注
	case 2:
		err := uc.repo.UnFollow(ctx, userId, toUserId)
		if err != nil {
			return errors.New("something wrong")
		}
	}
	return nil
}

func (uc *RelationUsecase) GetFollowNumber(ctx context.Context, userId uint32) (int64, error) {
	return uc.repo.GetFollowCount(ctx, userId)
}

func (uc *RelationUsecase) GetFollowerNumber(ctx context.Context, userId uint32) (int64, error) {
	return uc.repo.GetFollowerCount(ctx, userId)
}
