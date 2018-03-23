package main

import (
	"config"
	"flag"
	"log"
	"server"
)

var (
	// Choice 功能选项
	Choice int
)

func init() {

	config.Init()
	flag.IntVar(
		&Choice,
		"choice",
		3,
		"功能选择\n\t0:\t打印配置项\n\t1:\t使用配置进行数据下载\n\t2:\t报表-等待月时\n\t3:\t报表-公司角度")
	flag.Parse()
}

func main() {

	switch Choice {
	case 0:
		server.PrintConfig()
	case 1:
		server.Download()
	case 2:
		server.StatisticByMonth()
	case 3:
		server.StatisticByCompany()
	default:
		log.Fatal("UnSupport Choice")
	}
}
