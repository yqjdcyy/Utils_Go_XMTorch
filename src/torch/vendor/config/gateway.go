package config

import (
	"log"
	"time"
)

var (
	// Format 日期格式
	Format string
)

func init() {
	Format = "20060102"
}

// GatewayConfig  网关服务的配置
type GatewayConfig struct {
	c *ConfigFile

	From     time.Time
	To       time.Time
	Interval int
	Path     string

	UrlRentDetail string

	Cookie string

	StatisticPath string
}

func (g *GatewayConfig) init(c *ConfigFile) {
	g.c = c

	g.Interval = c.GetIntDefault("setting", "interval", 1)
	g.Path = c.GetStringDefault("setting", "path", "D:\\data\\export")

	to := c.GetStringDefault("setting", "to", "")
	if 0 == len(to) || 8 != len(to) {
		g.To = time.Now()
	} else {
		t, err := time.Parse(Format, to)
		if nil != err {
			log.Fatalf("fail to parse time.to[%s]", to)
		}
		g.To = t
	}

	from := c.GetStringDefault("setting", "from", "")
	if 0 == len(from) || 8 != len(from) {
		g.From = time.Now()
	} else {
		t, err := time.Parse(Format, from)
		if nil != err {
			log.Fatalf("fail to parse time.from[%s]", to)
		}
		g.From = t
	}

	g.UrlRentDetail= c.GetStringDefault("url", "rentDetail", "")
	if 0== len(g.UrlRentDetail){
		log.Fatalln("fail to request for empty url")
	}
	g.Cookie= c.GetStringDefault("auth", "cookie", "")
	if 0== len(g.Cookie){
		log.Fatalln("fail to request for empty cookie")
	}

	g.StatisticPath= c.GetStringDefault("statistic", "path", "")
}
