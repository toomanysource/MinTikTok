package data

import (
	pb "Atreus/api/user/service/v1"
	"Atreus/app/relation/service/internal/biz"
	"context"
	"errors"
	"google.golang.org/grpc"
)

type userRepo struct {
	client pb.UserServiceClient
}

func NewUserRepo(conn *grpc.ClientConn) UserRepo {
	return &userRepo{
		client: pb.NewUserServiceClient(conn),
	}
}

// GetUserInfos 接收User服务的回应，并转化为biz.User类型
func (u *userRepo) GetUserInfos(ctx context.Context, userIds []uint32) ([]*biz.User, error) {
	resp, err := u.client.GetUserInfos(ctx, &pb.UserInfosRequest{UserIds: userIds})
	if err != nil {
		return nil, err
	}

	// 判空
	if len(resp.Users) == 0 {
		return nil, errors.New("the user service did not search for any information")
	}

	users := make([]*biz.User, 0, len(resp.Users)+1)
	for _, user := range resp.Users {
		users = append(users, &biz.User{
			Id:              user.Id,
			Name:            user.Name,
			Avatar:          user.Avatar,
			BackgroundImage: user.BackgroundImage,
			Signature:       user.Signature,
			IsFollow:        user.IsFollow,
			FollowCount:     user.FollowCount,
			FollowerCount:   user.FollowerCount,
			TotalFavorite:   user.TotalFavorited,
			WorkCount:       user.WorkCount,
			FavoriteCount:   user.FavoriteCount,
		})
	}
	return users, nil
}

// UpdateFollow 接收User服务的回应
func (u *userRepo) UpdateFollow(ctx context.Context, userId uint32, followChange int32) error {
	resp, err := u.client.UpdateFollow(
		ctx, &pb.UpdateFollowRequest{UserId: userId, FollowChange: followChange})
	if err != nil {
		return err
	}
	if resp.StatusCode != 0 {
		return errors.New(resp.StatusMsg)
	}
	return nil
}

func (u *userRepo) UpdateFollower(ctx context.Context, userId uint32, followerChange int32) error {
	resp, err := u.client.UpdateFollower(
		ctx, &pb.UpdateFollowerRequest{UserId: userId, FollowerChange: followerChange})
	if err != nil {
		return err
	}
	if resp.StatusCode != 0 {
		return errors.New(resp.StatusMsg)
	}
	return nil
}
