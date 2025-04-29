package main

import (
	"web-starter/internal/config"
	"web-starter/internal/jobs"
	"web-starter/internal/router"
	"web-starter/internal/utils/db"
	"web-starter/internal/utils/glog"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy

func main() {
	config.InitConfig() // 配置
	glog.InitLogger()   // log
	db.InitDB()         // 连接数据库
	jobs.InitCronJob()  // 初始化定时任务
	router.InitServer() // 注册路由

}
