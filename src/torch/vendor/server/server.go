package server

import (
	"config"

	"bitbucket.org/ansenwork/ilog"
)

// Download 数据下载
func Download() {

	Circulate(config.Gateway.From, config.Gateway.To, config.Gateway.UrlRentDetail)
}

// PrintConfig 打印配置信息
func PrintConfig() {

	ilog.Infof("%v", config.Gateway.From)
	ilog.Infof("%v", config.Gateway.To)
	ilog.Infof("%v", config.Gateway.Path)
	ilog.Infof("%v", config.Gateway.Interval)
	ilog.Infof("%v", config.Gateway.UrlRentDetail)
	ilog.Infof("%v", config.Gateway.Cookie)
	ilog.Infof("%v", config.Gateway.StatisticPath)
}


// StatisticByMonth 申请时长角度
func StatisticByMonth(){

	CalcByMonth(config.Gateway.StatisticPath)
}
// StatisticByCompany 公司角度
func StatisticByCompany(){

	ilog.Info("TODO")
}