package main

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"hostinger-test/conf"
	"log"
)

func main() {
	// 初始化全局配置
	initGlobalConf()

	// 启动 goframe web框架
	gfServer()
}

func gfServer() {
	s := newGFServer()
	s.BindHandler("/health", healCheck)

	dbTest := s.Group("/db")
	dbTest.GET("/ping", dbPing)

	s.Run()
}

func newGFServer() *ghttp.Server {
	s := g.Server()
	s.SetPort(conf.GlobalC.Port.HTTPPort)
	if err := s.SetConfigWithMap(conf.GlobalC.GFServer.Config); err != nil {
		s.Logger().Fatalf(context.Background(), "gf server config error:%v", err)
	}
	return s
}

func initGlobalConf() {
	log.Println(conf.Init(context.Background(), "conf/sim.toml"))
}
