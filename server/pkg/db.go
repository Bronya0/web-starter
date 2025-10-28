package pkg

import (
	"log"
	"time"

	"gorm.io/datatypes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dsn = "host=127.0.0.1 port=5432 user=postgres dbname=postgres sslmode=disable TimeZone=Asia/Shanghai"
	DB  *gorm.DB
)

type Epub struct {
	Name       string
	Hash       string `gorm:"index"`
	Size       int64
	IP         string
	Addr       string
	Screen     datatypes.JSON `gorm:"type:jsonb"`
	Cpus       int
	IsUploaded bool `gorm:"column:is_uploaded"`
	gorm.Model
}

func NewDB() {
	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction:                   true, // 跳过默认事务
		PrepareStmt:                              true, // 缓存预编译
		DisableForeignKeyConstraintWhenMigrating: true, // 禁止迁移时创建外键
	})
	if err != nil {
		log.Fatal("Failed to connect to database!")
	}

	sqlDB, _ := db.DB()
	// 连接的最大数量
	sqlDB.SetMaxOpenConns(100)
	// 空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// 连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)
	// 连接在空闲状态下最长可保持的时间
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)

	err = db.AutoMigrate(&Epub{})
	if err != nil {
		log.Fatal(err)
	}

	DB = db
}
