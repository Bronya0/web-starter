package glog

import (
	"gin-starter/internal/config"
	"github.com/donnie4w/go-logger/logger"
	"log"
)

var (
	Log *logger.Logging
)

func InitLogger() {
	Log = initLogger(config.Conf.Log.Path)
}

// InitLogger pathFile: 日志全路径
func initLogger(path string) *logger.Logging {
	log.Println("初始化日志……" + path)
	l := logger.NewLogger()
	l.SetOption(&logger.Option{
		Level:     logger.LEVEL_INFO,
		Console:   true, // 控制台输出
		Format:    logger.FORMAT_LEVELFLAG | logger.FORMAT_SHORTFILENAME | logger.FORMAT_DATE | logger.FORMAT_MICROSECONDS,
		Formatter: "[{time}] {level} {file}: {message}\n",
		// size或者time模式
		FileOption: &logger.FileTimeMode{ // 这里用时间切割
			Filename:   path,            // 日志文件路径
			Timemode:   logger.MODE_DAY, // 按天
			Maxbuckup:  180,             // 最多备份日志文件数
			IsCompress: false,           // 是否压缩
		},
	})

	return l
}
