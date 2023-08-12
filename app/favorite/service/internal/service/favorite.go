package service

import (
	pb "Atreus/api/favorite/service/v1"
	"Atreus/app/favorite/service/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type FavoriteService struct {
	pb.UnimplementedFavoriteServiceServer
	fu  *biz.FavoriteUsecase
	log *log.Helper
}

func NewFavoriteService(fu *biz.FavoriteUsecase, logger log.Logger) *FavoriteService {
	return &FavoriteService{
		fu:  fu,
		log: log.NewHelper(log.With(logger, "model", "service/favorite")),
	}
}

func (s *FavoriteService) GetFavoriteList(
	ctx context.Context, req *pb.FavoriteListRequest) (*pb.FavoriteListReply, error) {
	reply := &pb.FavoriteListReply{StatusCode: 0, StatusMsg: "success"}
	videos, err := s.fu.GetFavoriteList(ctx, req.UserId, req.Token)
	if err != nil {
		reply.StatusCode = -1
		reply.StatusMsg = err.Error()
		return reply, nil
	}
	for _, video := range videos {
		reply.VideoList = append(reply.VideoList, &pb.Video{
			Id:    video.Id,
			Title: video.Title,
			Author: &pb.User{
				Id:              video.Author.Id,
				Name:            video.Author.Name,
				FollowCount:     video.Author.FollowCount,
				FollowerCount:   video.Author.FollowerCount,
				IsFollow:        video.Author.IsFollow,
				Avatar:          video.Author.Avatar,
				BackgroundImage: video.Author.BackgroundImage,
				Signature:       video.Author.Signature,
				TotalFavorited:  video.Author.TotalFavorited,
				WorkCount:       video.Author.WorkCount,
				FavoriteCount:   video.Author.FavoriteCount,
			},
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    video.IsFavorite,
		})
	}
	return reply, nil
}

func (s *FavoriteService) FavoriteAction(
	ctx context.Context, req *pb.FavoriteActionRequest) (*pb.FavoriteActionReply, error) {
	reply := &pb.FavoriteActionReply{StatusCode: 0, StatusMsg: "success"}
	err := s.fu.FavoriteAction(ctx, req.VideoId, req.ActionType, req.Token)
	if err != nil {
		reply.StatusCode = -1
		reply.StatusMsg = err.Error()
		return reply, nil
	}
	return reply, nil
}

func (s *FavoriteService) IsFavorite(
	ctx context.Context, req *pb.IsFavoriteRequest) (*pb.IsFavoriteReply, error) {
	reply := &pb.IsFavoriteReply{StatusCode: 0, StatusMsg: "success"}
	isFavorite, err := s.fu.IsFavorite(ctx, req.VideoId, req.UserId)
	if err != nil {
		reply.StatusCode = -1
		reply.StatusMsg = err.Error()
		return reply, nil
	}
	reply.IsFavorite = isFavorite
	return reply, nil
}
