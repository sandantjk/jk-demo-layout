package main

import (
	"jk-demo-layout/internal/conf"
	"jk-demo-layout/internal/fbase"
	"jk-demo-layout/internal/server"
	"log"
	"net/http"
)

func main() {
	config := &conf.Config{}
	err := config.LoadFile("./configs/application.yml")
	if err != nil {
		log.Printf("%+v\n", err)
		return
	}
	app := initApp(config, log.Default())
	if err := app.Start(); err != nil {
		log.Printf("%+v\n", err)
		return
	}
}

/*func initApp(config *conf.Config, logger *log.Logger) *fbase.App {
	userRepo := data.NewUserRepo()
	userBiz := biz.NewUserBiz(userRepo)
	userService := service.NewUserService(userBiz, logger)
	server0 := server.NewHttpServer0(config, userService)
	server1 := server.NewHttpServer1(config, userService)
	return newApp(logger, server0, server1)
}*/

func newApp(logger *log.Logger, server0 *server.HttpServer0, server1 *server.HttpServer1) *fbase.App {
	return fbase.New(fbase.Name("demo-layout"), fbase.Logger(logger), fbase.Servers((*http.Server)(server0), (*http.Server)(server1)))
}
