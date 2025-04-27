package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"web-starter/internal/utils/glog"
)

type PgSql struct {
	gormConfig *gorm.Config
}

func (p *PgSql) NewDB(dsn string) *gorm.DB {

	pgsqlConfig := postgres.Config{
		DSN:                  dsn, // DSN data source name
		PreferSimpleProtocol: false,
	}
	db, err := gorm.Open(postgres.New(pgsqlConfig), p.gormConfig)

	if err != nil {
		glog.Log.Error(err)
		panic(err)
	} else {
		glog.Log.Info("数据库连接成功...")
	}

	return db

}
