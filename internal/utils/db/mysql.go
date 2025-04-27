package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"web-starter/internal/utils/glog"
)

type Mysql struct {
	gormConfig *gorm.Config
}

func (m *Mysql) NewDB(dsn string) *gorm.DB {

	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	db, err := gorm.Open(mysql.New(mysqlConfig), m.gormConfig)
	if err != nil {
		glog.Log.Error(err)
		panic(err)
	} else {
		glog.Log.Info("数据库连接成功...")
	}
	db.InstanceSet("gorm:table_options", "ENGINE=innodb")

	return db
}
