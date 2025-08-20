package glog

import (
	"web-starter/internal/config"

	logging "github.com/donnie4w/go-logger/logger"
)

var (
	errLogger *logging.Logging
	Logger    *logging.Logging
)

func InitLogger() {
	errLogger = InitErrLogger(config.Conf.Log.Err)
	Logger = InitAllLogger(config.Conf.Log.Path)
}

// InitAllLogger 全部日志
func InitAllLogger(path string) *logging.Logging {
	logger := logging.NewLogger()
	logger.SetOption(&logging.Option{
		Level:     logging.LEVEL_INFO,
		Console:   true, // 控制台输出
		Format:    logging.FORMAT_LEVELFLAG | logging.FORMAT_SHORTFILENAME | logging.FORMAT_DATE | logging.FORMAT_MICROSECONDS,
		Formatter: "[{time}] {level} {file}: {message}\n",
		// size或者time模式
		FileOption: &logging.FileTimeMode{ // 这里用时间切割
			Filename:   path,             // 日志文件路径
			Timemode:   logging.MODE_DAY, // 按天
			Maxbuckup:  7,                // 最多备份日志文件数
			IsCompress: false,            // 是否压缩
		},
		// 自定义处理函数，err及以上级别日志额外写入err.log
		CustomHandler: func(lc *logging.LogContext) bool {
			if errLogger != nil {
				if lc.Level == logging.LEVEL_ERROR {
					errLogger.Error(lc.Args...)
				}
				if lc.Level == logging.LEVEL_FATAL {
					errLogger.Fatal(lc.Args...)
				}
			}
			return true
		},
	},
	)
	return logger
}

// InitErrLogger 错误日志
func InitErrLogger(path string) *logging.Logging {
	logger := logging.NewLogger()
	logger.SetOption(&logging.Option{
		Level:     logging.LEVEL_ERROR,
		Console:   false, // 控制台输出
		Format:    logging.FORMAT_LEVELFLAG | logging.FORMAT_SHORTFILENAME | logging.FORMAT_DATE | logging.FORMAT_MICROSECONDS,
		Formatter: "[{time}] {level} {file}: {message}\n",
		FileOption: &logging.FileTimeMode{ // 这里用时间切割
			Filename:   path,             // 日志文件路径
			Timemode:   logging.MODE_DAY, // 按天
			Maxbuckup:  7,                // 最多备份日志文件数
			IsCompress: false,            // 是否压缩
		},
	},
	)
	return logger
}
