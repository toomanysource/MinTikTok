package middleware

import (
	"context"

	"Atreus/pkg/errorX"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/golang-jwt/jwt/v4"
)

// TokenParseAll 所有用户都可以访问，非必须token
func TokenParseAll(keyFunc jwt.Keyfunc) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				if hr, ok := tr.(*http.Transport); ok {
					values := hr.Request().URL.Query()
					tokenString := values.Get("token")
					if tokenString == "" {
						ctx = context.WithValue(ctx, "user_id", uint32(0))
						return handler(ctx, req)
					}
					token, err := jwt.Parse(tokenString, keyFunc)
					if err != nil {
						return nil, errorX.New(-1, err.Error())
					}
					if !token.Valid {
						return nil, errorX.New(-1, "token is invalid")
					}
					ctx = context.WithValue(ctx, "user_id", uint32(token.Claims.(jwt.MapClaims)["user_id"].(float64)))
				}
			}
			return handler(ctx, req)
		}
	}
}
