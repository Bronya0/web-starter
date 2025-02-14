package jobs

import (
	"fmt"
	"gin-starter/internal/config"
	"gin-starter/internal/model"
	"gin-starter/internal/util/db"
	"gin-starter/internal/util/glog"
	"github.com/go-co-op/gocron/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"reflect"
	"runtime"
	"time"
)

func InitCronJob() {
	s := initScheduler()
	// 初始化表
	if config.Conf.DB.Enable {
		model.AutoMigrateJobs()
	}
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

	if config.Conf.DB.Enable {
		saveOrUpdate(jobName, crontab, function)
	}

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

// saveOrUpdate 保存job or 更新
func saveOrUpdate(jobName, crontab string, fun any) {
	// 如果任务存在(jobName+fun)，则更新crontab，并把其他字段清空；否则新建
	if db.DB.Where("name = ? and func = ?", jobName, getFunctionName(fun)).First(&model.CronJob{}).RowsAffected > 0 {
		// 更新crontab
		db.DB.Model(&model.CronJob{}).Where("name = ? and func = ?", jobName, getFunctionName(fun)).
			UpdateColumns(map[string]interface{}{
				"crontab":        crontab,
				"last_run_start": nil,
				"last_run_end":   nil,
				"run_count":      0,
				"success":        true,
				"error":          nil,
			})
		return
	}

	// 记录任务到数据库
	cronJob := &model.CronJob{
		Name:    jobName,
		Crontab: crontab,
		Func:    getFunctionName(fun),
	}
	result := db.DB.Where("name = ?", jobName).Create(cronJob)
	if result.Error != nil {
		glog.Log.Errorf("任务记录创建失败: %v", result.Error)
		panic(result.Error)
	}
}

// beforeListener  任务运行前执行
func beforeListener() gocron.EventListener {
	return gocron.BeforeJobRuns(func(jobID uuid.UUID, jobName string) {
		glog.Log.Infof("Job %s is start running...", jobName)

		if !config.Conf.DB.Enable {
			return
		}
		// 更新任务信息
		db.DB.Model(&model.CronJob{}).Where("name = ?", jobName).
			UpdateColumns(map[string]interface{}{
				"last_run_start": time.Now(),
			})
	})
}

// afterListener  用于监听作业何时运行且没有错误，然后运行提供的函数。
func afterListener() gocron.EventListener {
	return gocron.AfterJobRuns(func(jobID uuid.UUID, jobName string) {
		glog.Log.Infof("Job %s is running end", jobName)

		if !config.Conf.DB.Enable {
			return
		}
		// 更新任务信息
		db.DB.Model(&model.CronJob{}).Where("name = ?", jobName).
			UpdateColumns(map[string]interface{}{
				"last_run_end": time.Now(),
				"run_count":    gorm.Expr("run_count + 1"),
				"success":      true,
			})
	})
}

// panicListener  panic监听器
func panicListener() gocron.EventListener {
	return gocron.AfterJobRunsWithPanic(func(jobID uuid.UUID, jobName string, recoverData any) {
		glog.Log.Errorf("Job Panic！！！：jobName: %s jobID: (%s): %+v\n", jobName, jobID, recoverData)

		if !config.Conf.DB.Enable {
			return
		}
		// 更新任务信息
		db.DB.Model(&model.CronJob{}).Where("id = ?", jobID).
			UpdateColumns(map[string]interface{}{
				"last_run_end": time.Now(),
				"run_count":    gorm.Expr("run_count + 1"),
				"success":      false,
				"error":        fmt.Sprintf("%+v", recoverData),
			})
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
