// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"jk-demo-layout/internal/biz"
	"jk-demo-layout/internal/conf"
	"jk-demo-layout/internal/data"
	"jk-demo-layout/internal/fbase"
	"jk-demo-layout/internal/server"
	"jk-demo-layout/internal/service"
	"log"
)

// Injectors from wire.go:

func initApp(config *conf.Config, logger *log.Logger) *fbase.App {
	userRepo := data.NewUserRepo()
	userBiz := biz.NewUserBiz(userRepo)
	userService := service.NewUserService(userBiz, logger)
	httpServer0 := server.NewHttpServer0(config, userService)
	httpServer1 := server.NewHttpServer1(config, userService)
	app := newApp(logger, httpServer0, httpServer1)
	return app
}
