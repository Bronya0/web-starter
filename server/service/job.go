package service

import (
	"reflect"
	"runtime"
	"time"
	"web-starter/server/pkg"

	"github.com/go-co-op/gocron/v2"
	"github.com/google/uuid"
)

func InitCronJob() {
	s := initScheduler()
	addJobs(s)
	start(s)
}

// addJobs 添加任务
func addJobs(s *gocron.Scheduler) {
	// 添加一个每10秒执行一次的任务
	//addJob(s, "PingUrl", "*/10 * * * * *", Ping)

}

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
		gocron.WithEventListeners(panicListener(), beforeListener(), afterListener()),
		gocron.WithName(jobName),
		gocron.WithTags(jobName, getFunctionName(function)),  // 用于删除
		gocron.WithSingletonMode(gocron.LimitModeReschedule), // 避免重叠运行
	)

	if err != nil {
		pkg.Logger.Errorf("定时任务注册失败！: %v：%v", jobName, err)
		panic(err)
	}
	pkg.Logger.Infof("定时任务: %v 注册成功", jobName)
}

func initScheduler() *gocron.Scheduler {
	s, err := gocron.NewScheduler(
		gocron.WithLocation(time.Local),
		gocron.WithGlobalJobOptions(),
	)
	if err != nil {
		pkg.Logger.Errorf("initScheduler失败！: %v", err)
		panic("initScheduler失败！: " + err.Error())
	}
	return &s
}

func start(s *gocron.Scheduler) {
	scheduler := *s
	scheduler.Start()
	pkg.Logger.Info("定时任务启动成功...")
}

// beforeListener  任务运行前执行
func beforeListener() gocron.EventListener {
	return gocron.BeforeJobRuns(func(jobID uuid.UUID, jobName string) {
		pkg.Logger.Infof("Job %s is start running...", jobName)
	})
}

// afterListener  用于监听作业何时运行且没有错误，然后运行提供的函数。
func afterListener() gocron.EventListener {
	return gocron.AfterJobRuns(func(jobID uuid.UUID, jobName string) {
		pkg.Logger.Infof("Job %s is running end", jobName)
	})
}

// panicListener  panic监听器
func panicListener() gocron.EventListener {
	return gocron.AfterJobRunsWithPanic(func(jobID uuid.UUID, jobName string, recoverData any) {
		pkg.Logger.Errorf("Job Panic！！！：jobName: %s jobID: (%s): %+v\n", jobName, jobID, recoverData)
	})
}

// getFunctionName 获取函数名称
func getFunctionName(fun any) string {
	// 使用反射获取函数的值
	value := reflect.ValueOf(fun)
	if value.Kind() != reflect.Func {
		return "unknown"
	}
	// 获取函数的指针
	pc := value.Pointer()
	if pc == 0 {
		return "unknown"
	}
	f := runtime.FuncForPC(pc)
	if f == nil {
		return "unknown"
	}
	return f.Name()
}
