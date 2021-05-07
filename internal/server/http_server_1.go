package server

import (
	"jk-demo-layout/api/demo"
	"jk-demo-layout/internal/conf"
	"jk-demo-layout/internal/service"
	"net/http"
	"strconv"
)

type HttpServer1 http.Server

func NewHttpServer1(config *conf.Config, user *service.UserService) *HttpServer1 {
	mux := http.NewServeMux()
	mux.HandleFunc("/", demo.NewHelloHandler(user))
	//
	return &HttpServer1{
		Addr:    ":" + strconv.Itoa(config.Server1Port),
		Handler: mux,
	}
}
