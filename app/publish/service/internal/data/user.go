package data

import (
	pb "Atreus/api/user/service/v1"
	"Atreus/app/publish/service/internal/biz"
	"Atreus/app/publish/service/internal/server"
	"context"
	"errors"
)

type userRepo struct {
	client pb.UserServiceClient
}

func NewUserRepo(conn server.UserConn) UserRepo {
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
			ID:              user.Id,
			Name:            user.Name,
			Avatar:          user.Avatar,
			BackgroundImage: user.BackgroundImage,
			Signature:       user.Signature,
			IsFollow:        user.IsFollow,
			FollowCount:     user.FollowCount,
			FollowerCount:   user.FollowerCount,
			TotalFavorited:  user.TotalFavorited,
			WorkCount:       user.WorkCount,
			FavoriteCount:   user.FavoriteCount,
		})
	}
	return users, nil
}
