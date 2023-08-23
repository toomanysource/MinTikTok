// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"Atreus/app/message/service/internal/biz"
	"Atreus/app/message/service/internal/conf"
	"Atreus/app/message/service/internal/data"
	"Atreus/app/message/service/internal/server"
	"Atreus/app/message/service/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, jwt *conf.JWT, logger log.Logger) (*kratos.App, func(), error) {
	db := data.NewMysqlConn(confData)
	kafkaConn := data.NewKafkaConn(confData)
	client := data.NewRedisConn(confData, logger)
	dataData, cleanup, err := data.NewData(db, kafkaConn, client, logger)
	if err != nil {
		return nil, nil, err
	}
	messageRepo := data.NewMessageRepo(dataData, logger)
	messageUsecase := biz.NewMessageUsecase(messageRepo, jwt, logger)
	messageService := service.NewMessageService(messageUsecase, logger)
	grpcServer := server.NewGRPCServer(confServer, messageService, logger)
	httpServer := server.NewHTTPServer(confServer, messageService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
