package service

import (
	pb "Atreus/api/user/service/v1"
	"context"
	"google.golang.org/grpc"
)

type UserServiceClient struct {
	uc   pb.UserServiceClient
	conn *grpc.ClientConn
}

func NewUserServiceClient(conn *grpc.ClientConn) *UserServiceClient {
	return &UserServiceClient{
		conn: conn,
		uc:   pb.NewUserServiceClient(conn),
	}
}

func (c *UserServiceClient) GetUserInfoByUserIds(
	ctx context.Context, userIds []uint32) (*pb.ClientUserInfosReply, error) {
	return c.uc.GetUserInfoByUserIds(ctx, &pb.ClientUserInfoByUserIdsRequest{UserIds: userIds})
}
