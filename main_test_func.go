package main

import (
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/net/ghttp"
)

func healCheck(r *ghttp.Request) {
	r.Response.Write("health ok!")
}

func dbPing(r *ghttp.Request) {
	db, err := gdb.Instance()
	if err != nil {
		r.Response.Write("db instance error!", err)
		return
	}
	err = db.PingMaster()
	if err != nil {
		r.Response.Write("db ping error!", err)
		return
	}
	r.Response.Write("db test!")
}
