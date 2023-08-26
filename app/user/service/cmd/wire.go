//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/toomanysource/atreus/app/user/service/internal/biz"
	"github.com/toomanysource/atreus/app/user/service/internal/conf"
	"github.com/toomanysource/atreus/app/user/service/internal/data"
	"github.com/toomanysource/atreus/app/user/service/internal/server"
	"github.com/toomanysource/atreus/app/user/service/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.Client, *conf.JWT, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
