package model

import (
	"time"
	"web-starter/internal/utils/db"
	"web-starter/internal/utils/glog"
)

type CronJob struct {
	ID           int32      `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name         string     `gorm:"type:varchar(256)"`
	Crontab      string     `gorm:"type:varchar(256)"`
	Func         string     `gorm:"type:varchar(1024)"`
	LastRunStart *time.Time `gorm:"type:timestamp"`
	LastRunEnd   *time.Time `gorm:"type:timestamp"`
	RunCount     int64
	Success      bool       // 执行是否成功
	Error        string     `gorm:"type:text"` // 错误信息
	CreateTime   *time.Time `gorm:"column:create_time;type:TIMESTAMP;not null;default:CURRENT_TIMESTAMP;" json:"create_time"`
	UpdateTime   *time.Time `gorm:"column:update_time;type:TIMESTAMP;not null;default:CURRENT_TIMESTAMP;" json:"update_time"`
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
	glog.Log.Info("自动迁移成功")
}
