package model

import (
	"time"
	"web-starter/internal/utils/db"
	"web-starter/internal/utils/glog"
)

type CronJob struct {
	BaseModel
	Name         string     `gorm:"type:varchar(256)"`
	Crontab      string     `gorm:"type:varchar(256)"`
	Func         string     `gorm:"type:varchar(1024)"`
	LastRunStart *time.Time `gorm:"type:timestamp"`
	LastRunEnd   *time.Time `gorm:"type:timestamp"`
	RunCount     int64
	Success      bool   // 执行是否成功
	Error        string `gorm:"type:text"` // 错误信息
}

func (m *CronJob) TableName() string {
	return "public.cron_job" // return "schema.table"
}

// AutoMigrateJobs 自动迁移
func AutoMigrateJobs() {
	err := db.DB.AutoMigrate(&CronJob{})
	if err != nil {
		panic("自动迁移失败: " + err.Error())
	}
	glog.Logger.Info("自动迁移成功")
}
