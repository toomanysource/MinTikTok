package server

import (
	v1 "Atreus/api/comment/service/v1"
	"Atreus/app/comment/service/internal/conf"
	"Atreus/app/comment/service/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, greeter *service.CommentService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			logging.Server(logger),
			// 服务间数据传输的jwt校验，如需使用可解除注释
			//jwt.Server(func(token *jwtv4.Token) (any, error) {
			//	return []byte(j.Grpc.TokenKey), nil
			//}),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}

	srv := grpc.NewServer(opts...)
	v1.RegisterCommentServiceServer(srv, greeter)
	return srv
}
