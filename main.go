package main

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"hostinger-test/conf"
	"log"
	"net/http"
)

func main() {
	// 初始化全局配置
	initGlobalConf()

	// 启动 goframe web框架
	gfServer()
	select {}
}

func originHTTPServer() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("health ok 3"))
		log.Printf("health check, err:%v", err)
	})
	http.ListenAndServe(":9002", nil)
}

func gfServer() {
	s := g.Server()
	s.BindHandler("/health", func(r *ghttp.Request) {
		r.Response.Write("health ok!")
	})

	s.SetPort(conf.GlobalC.Port.HTTPPort)
	s.Run()
}

func initGlobalConf() {
	log.Println(conf.Init(context.Background(), "conf/sim.toml"))
}
