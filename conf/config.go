package conf

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gcfg"
)

var (
	GlobalC = C{} // 全局配置，只允许初始化一次
)

type C struct {
	Port      Port
	GFServer  GFServer
	DefaultDB DefaultDB
}

func (c C) String() string {
	str, _ := json.Marshal(c)
	return string(str)
}

type Port struct {
	HTTPPort int
}

func (p *Port) init(ctx context.Context, cfgInstance *gcfg.Config) error {
	httpPort, err := cfgInstance.Get(ctx, "port.http")
	if err != nil {
		return err
	}
	p.HTTPPort = httpPort.Int()
	return nil
}

type GFServer struct {
	Config map[string]any
}

func (gfs *GFServer) init(ctx context.Context, cfgInstance *gcfg.Config) error {
	serverConf, err := cfgInstance.Get(ctx, "server")
	if err != nil {
		return err
	}
	gfs.Config = serverConf.Map()
	return nil
}

type DefaultDB struct {
}

func (gfd *DefaultDB) init(ctx context.Context, cfgInstance *gcfg.Config) error {
	// GF 有自己的注册driverMap
	dbConf, err := cfgInstance.Get(ctx, "database.default")
	if err != nil {
		return err
	}
	// 注册mysql驱动
	if err = gdb.Register("mysql", &gdb.DriverDefault{}); err != nil {
		return err
	}

	var gdbConfItems gdb.ConfigNode
	if err = dbConf.Structs(&gdbConfItems); err != nil {
		return err
	}
	// 初始化配置
	if err = gdb.AddConfigNode("default", gdbConfItems); err != nil {
		return err
	}
	return nil
}

func Init(ctx context.Context, confPath string) error {
	cfgInstance := gcfg.Instance(confPath)

	// http 配置
	if err := GlobalC.Port.init(ctx, cfgInstance); err != nil {
		return err
	}

	// gf server 配置
	if err := GlobalC.GFServer.init(ctx, cfgInstance); err != nil {
		return err
	}

	// 默认数据库配置
	if err := GlobalC.DefaultDB.init(ctx, cfgInstance); err != nil {
		return err
	}

	return nil
}
