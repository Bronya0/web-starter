package model

import (
	"time"
)

type BaseModel struct {
	ID         uint      `gorm:"primarykey" json:"id"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime" json:"update_time"`
}
