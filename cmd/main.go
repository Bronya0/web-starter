package main

import (
	"web-starter/internal/config"
	"web-starter/internal/jobs"
	"web-starter/internal/router"
	"web-starter/internal/utils/db"
	"web-starter/internal/utils/glog"
)

func main() {
	config.InitConfig() // 配置
	glog.InitLogger()   // log
	db.InitDB()         // 连接数据库
	jobs.InitCronJob()  // 初始化定时任务
	router.InitServer() // 注册路由

}
