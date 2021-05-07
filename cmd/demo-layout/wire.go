//+build wireinject

package main

import (
	"github.com/google/wire"
	"jk-demo-layout/internal/biz"
	"jk-demo-layout/internal/conf"
	"jk-demo-layout/internal/data"
	"jk-demo-layout/internal/fbase"
	"jk-demo-layout/internal/server"
	"jk-demo-layout/internal/service"
	"log"
)

func initApp(config *conf.Config, logger *log.Logger) *fbase.App {
	wire.Build(server.WireSet, data.WireSet, biz.WireSet, service.WireSet, newApp)
	return &fbase.App{}
}
