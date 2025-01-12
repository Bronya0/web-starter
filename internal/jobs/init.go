package jobs

import (
	"gin-starter/internal/util/glog"
	"github.com/go-co-op/gocron/v2"
	"github.com/google/uuid"
	"time"
)

func InitCronJob() {
	// 创建一个新的调度器
	s := initScheduler()

	// 添加一个每10秒执行一次的任务
	//addJob(s, "print1", "*/10 * * * * *", PrintJob)

	// 启动调度器
	start(s)
}

// 通用job构造函数
func addJob(s *gocron.Scheduler, jobName string, crontab string, function any, parameters ...any) {
	// 超时请自己在任务中处理，不在外面做。
	scheduler := *s
	_, err := scheduler.NewJob(
		gocron.CronJob(
			crontab,
			true,
		),
		gocron.NewTask(
			function,
			parameters...,
		),
		gocron.WithEventListeners(jobRecover()),
		gocron.WithName(jobName),
		gocron.WithTags(jobName),                             // 用于s删除
		gocron.WithSingletonMode(gocron.LimitModeReschedule), // 避免重叠运行
	)

	if err != nil {
		glog.Log.Errorf("定时任务注册失败！: %v：%v", jobName, err)
		panic(err)
	}
	glog.Log.Infof("定时任务: %v 注册成功", jobName)
}

func initScheduler() *gocron.Scheduler {
	s, err := gocron.NewScheduler(
		gocron.WithLocation(time.Local),
		gocron.WithGlobalJobOptions(),
	)
	if err != nil {
		glog.Log.Errorf("initScheduler失败！: %v", err)
		panic("initScheduler失败！: " + err.Error())
	}
	return &s
}

func start(s *gocron.Scheduler) {
	scheduler := *s
	scheduler.Start()
	glog.Log.Info("定时任务启动成功...")
}

func jobRecover() gocron.EventListener {
	return gocron.AfterJobRunsWithPanic(func(jobID uuid.UUID, jobName string, recoverData any) {
		glog.Log.Errorf("Job Panic！！！：jobName: %s jobID: (%s): %v\n", jobName, jobID, recoverData)
	})
}

func PrintJob() {
	glog.Log.Info("主人你好,定时任务运行...")
	panic("定时任务运行...")
}
