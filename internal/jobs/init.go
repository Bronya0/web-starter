package jobs

import (
	"gin-starter/internal/util/glog"
	"github.com/robfig/cron/v3"
	"log"
)

func InitCronJob() {
	// 自定义日志记录器
	logger := cron.VerbosePrintfLogger(log.New(log.Writer(), "cron: ", log.LstdFlags))

	c := cron.New(
		cron.WithLogger(logger),
		cron.WithChain(
			cron.Recover(logger), // 捕获 panic 并记录日志
		),
	)

	// 依次是 分 时 日 月 周。非常自由
	// @every 1s、@every 1h、@every 1m、@every 1m2s、@every 1h30m10s
	_, err := c.AddFunc("@every 1m", PrintJob)
	if err != nil {
		panic(err)
	}
	c.Start()

	glog.Log.Info("定时任务加载成功...")

}

func PrintJob() {
	glog.Log.Info("主人你好,定时任务运行...")
	panic("定时任务运行...")
}
