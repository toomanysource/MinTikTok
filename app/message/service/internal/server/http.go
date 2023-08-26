package server

import (
	v1 "Atreus/api/message/service/v1"
	"Atreus/app/message/service/internal/conf"
	"Atreus/app/message/service/internal/service"
	"Atreus/middleware"
	"Atreus/pkg/errorX"

	"github.com/go-kratos/kratos/v2/middleware/validate"

	"github.com/golang-jwt/jwt/v4"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, t *conf.JWT, greeter *service.MessageService, logger log.Logger) *http.Server {
	opts := []http.ServerOption{
		http.ErrorEncoder(errorX.ErrorEncoder),
		http.Middleware(

			validate.Validator(),
			middleware.TokenParseAll(func(token *jwt.Token) (interface{}, error) {
				return []byte(t.Http.TokenKey), nil
			}),
			recovery.Recovery(),
			logging.Server(logger),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterMessageServiceHTTPServer(srv, greeter)
	return srv
}
