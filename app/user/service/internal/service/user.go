package service

import (
	"Atreus/app/user/service/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "Atreus/api/user/service/v1"
)

type UserService struct {
	pb.UnimplementedUserServiceServer

	log *log.Helper

	uc *biz.UserUsecase
}

func NewUserService(uc *biz.UserUsecase, logger log.Logger) *UserService {
	return &UserService{uc: uc, log: log.NewHelper(logger)}
}

func (s *UserService) UserRegister(ctx context.Context, req *pb.UserRegisterRequest) (*pb.UserRegisterReply, error) {
	user, err := s.uc.Register(context.TODO(), req.Username, req.Password)
	if err != nil {
		return &pb.UserRegisterReply{
			StatusCode: 300,
			StatusMsg:  err.Error(),
		}, nil
	}
	return &pb.UserRegisterReply{
		StatusCode: 0,
		StatusMsg:  "success",
		UserId:     user.Id,
		Token:      user.Token,
	}, nil
}

func (s *UserService) UserLogin(ctx context.Context, req *pb.UserLoginRequest) (*pb.UserLoginReply, error) {
	user, err := s.uc.Login(context.TODO(), req.Username, req.Password)
	if err != nil {
		return &pb.UserLoginReply{
			StatusCode: 300,
			StatusMsg:  err.Error(),
		}, nil
	}
	return &pb.UserLoginReply{
		StatusCode: 0,
		StatusMsg:  "success",
		UserId:     user.Id,
		Token:      user.Token,
	}, nil
}

func (s *UserService) GetUserInfo(ctx context.Context, req *pb.UserInfoRequest) (*pb.UserInfoReply, error) {
	user, err := s.uc.GetInfo(context.TODO(), req.UserId, req.Token)
	if err != nil {
		return &pb.UserInfoReply{
			StatusCode: 300,
			StatusMsg:  err.Error(),
		}, nil
	}
	reply := &pb.UserInfoReply{
		StatusCode: 0,
		StatusMsg:  "success",
	}
	// no need IsFollow
	reply.User = &pb.User{
		Id:              user.Id,
		Name:            user.Name,
		FollowCount:     user.FollowCount,
		FollowerCount:   user.FollowerCount,
		Avatar:          user.Avatar,
		BackgroundImage: user.BackgroundImage,
		Signature:       user.Signature,
		TotalFavorited:  user.TotalFavorited,
		WorkCount:       user.WorkCount,
		FavoriteCount:   user.FavoriteCount,
	}
	return reply, nil
}

func (s *UserService) UpdateUserInfo(ctx context.Context, req *pb.UpdateUserInfoRequest) (*pb.UpdateUserInfoReply, error) {
	info := &biz.UserInfo{
		Name:            req.Name,
		Avatar:          req.Avatar,
		BackgroundImage: req.BackgroundImage,
		Signature:       req.Signature,
	}

	err := s.uc.UpdateInfo(context.TODO(), info)
	if err != nil {
		return &pb.UpdateUserInfoReply{
			StatusCode: 300,
			StatusMsg:  err.Error(),
		}, nil
	}
	return &pb.UpdateUserInfoReply{
		StatusCode: 0,
		StatusMsg:  "success",
	}, nil
}

func (s *UserService) GetUserInfos(ctx context.Context, req *pb.UserInfosRequest) (*pb.UserInfosReply, error) {
	users, err := s.uc.GetInfos(ctx, req.UserId, req.UserIds)
	if err != nil {
		return nil, err
	}
	reply := &pb.UserInfosReply{}
	reply.Users = make([]*pb.User, len(users))
	for i, user := range users {
		reply.Users[i] = &pb.User{
			Id:              user.Id,
			Name:            user.Name,
			FollowCount:     user.FollowCount,
			FollowerCount:   user.FollowerCount,
			Avatar:          user.Avatar,
			BackgroundImage: user.BackgroundImage,
			Signature:       user.Signature,
			TotalFavorited:  user.TotalFavorited,
			IsFollow:        user.IsFollow,
			WorkCount:       user.WorkCount,
			FavoriteCount:   user.FavoriteCount,
		}
	}
	return reply, nil
}

func (s *UserService) UpdateFollow(ctx context.Context, req *pb.UpdateFollowRequest) (*emptypb.Empty, error) {
	err := s.uc.UpdateFollow(ctx, req.UserId, req.FollowChange)
	return &emptypb.Empty{}, err
}

func (s *UserService) UpdateFollower(ctx context.Context, req *pb.UpdateFollowerRequest) (*emptypb.Empty, error) {
	err := s.uc.UpdateFollower(ctx, req.UserId, req.FollowerChange)
	return &emptypb.Empty{}, err
}

func (s *UserService) UpdateFavorited(ctx context.Context, req *pb.UpdateFavoritedRequest) (*emptypb.Empty, error) {
	err := s.uc.UpdateFavorited(ctx, req.UserId, req.FavoritedChange)
	return &emptypb.Empty{}, err
}

func (s *UserService) UpdateWork(ctx context.Context, req *pb.UpdateWorkRequest) (*emptypb.Empty, error) {
	err := s.uc.UpdateWork(ctx, req.UserId, req.WorkChange)
	return &emptypb.Empty{}, err
}

func (s *UserService) UpdateFavorite(ctx context.Context, req *pb.UpdateFavoriteRequest) (*emptypb.Empty, error) {
	err := s.uc.UpdateFavorite(ctx, req.UserId, req.FavoriteChange)
	return &emptypb.Empty{}, err
}
