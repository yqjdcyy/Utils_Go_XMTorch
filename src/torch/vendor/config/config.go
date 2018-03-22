package config

import (
	"flag"

	"bitbucket.org/ansenwork/ilog"
)

// Gateway 网关服务配置
var Gateway *GatewayConfig
var gatewayPath *string

func init() {
	gatewayPath = flag.String("c", "D:\\work\\git\\yao\\go\\Utils_Go_XMTorch\\src\\torch\\properties\\gateway.properties", "通用系统配置文件目录")
}

// Init 使用默认配置进行初始化
func Init() {

	createGateway(*gatewayPath)
}

func createGateway(path string) {
	var g = &GatewayConfig{}

	c, e := ReadConfigFile(path)
	if e != nil {
		ilog.Panicf("create Gateway config panic:%v", e)
		panic(e)
	}

	g.init(c)

	Gateway = g
}
