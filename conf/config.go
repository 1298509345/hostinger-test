package conf

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	GlobalC = C{} // 全局配置，只允许初始化一次
)

type C struct {
	Port     Port
	Database Database
	Logger   GFServerLogger
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

type Database struct {
	Dsn     string
	MaxIdle int
}

func (d *Database) init(ctx context.Context, cfgInstance *gcfg.Config) error {
	dsn, err := cfgInstance.Get(ctx, "database.dsn")
	if err != nil {
		return err
	}
	d.Dsn = dsn.String()
	maxIdle, err := cfgInstance.Get(ctx, "database.max_idle")
	if err != nil {
		return err
	}
	d.MaxIdle = maxIdle.Int()
	return nil
}

func Init(ctx context.Context, confPath string) error {
	cfgInstance := gcfg.Instance(confPath)

	// http 配置
	if err := GlobalC.Port.init(ctx, cfgInstance); err != nil {
		return err
	}

	// 数据库配置
	if err := GlobalC.Database.init(ctx, cfgInstance); err != nil {
		return err
	}

	return nil
}
