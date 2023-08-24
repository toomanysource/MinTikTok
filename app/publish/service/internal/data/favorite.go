package data

import (
	"context"

	pb "Atreus/api/favorite/service/v1"
	"Atreus/app/publish/service/internal/server"
)

type favoriteRepo struct {
	client pb.FavoriteServiceClient
}

func NewFavoriteRepo(conn server.FavoriteConn) FavoriteRepo {
	return &favoriteRepo{
		client: pb.NewFavoriteServiceClient(conn),
	}
}

// IsFavorite 接收favorite服务的回应
func (u *favoriteRepo) IsFavorite(ctx context.Context, userId uint32, videoIds []uint32) ([]bool, error) {
	resp, err := u.client.IsFavorite(ctx, &pb.IsFavoriteRequest{UserId: userId, VideoIds: videoIds})
	if err != nil {
		return nil, err
	}
	return resp.IsFavorite, nil
}
