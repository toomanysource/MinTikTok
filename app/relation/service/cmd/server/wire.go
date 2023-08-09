//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"Atreus/app/relation/service/internal/biz"
	"Atreus/app/relation/service/internal/conf"
	"Atreus/app/relation/service/internal/data"
	"Atreus/app/relation/service/internal/server"
	"Atreus/app/relation/service/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}