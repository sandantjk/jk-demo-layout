package server

import (
	"jk-demo-layout/api/demo"
	"jk-demo-layout/internal/conf"
	"jk-demo-layout/internal/service"
	"net/http"
	"strconv"
)

type HttpServer0 http.Server

func NewHttpServer0(config *conf.Config, user *service.UserService) *HttpServer0 {
	mux := http.NewServeMux()
	mux.HandleFunc("/", demo.NewHelloHandler(user))
	//
	return &HttpServer0{
		Addr:    ":" + strconv.Itoa(config.Server0Port),
		Handler: mux,
	}
}
