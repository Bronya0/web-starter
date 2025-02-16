package main

import (
	"gin-starter/internal/config"
	"gin-starter/internal/jobs"
	"gin-starter/internal/router"
	"gin-starter/internal/utils/db"
	"gin-starter/internal/utils/glog"
	"gin-starter/internal/utils/trans"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy

func main() {
	config.InitConfig()   // 配置
	glog.InitLogger()     // log
	db.InitDB()           // 连接数据库
	jobs.InitCronJob()    // 初始化定时任务
	trans.InitTrans("zh") // 初始化校验器，并本地化，zh/en
	router.InitServer()   // 注册路由，启动gin服务

}
