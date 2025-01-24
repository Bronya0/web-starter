package config

import (
	"fmt"
	"gin-starter/internal/common"
	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"path/filepath"
)

/*
只允许依赖common包
*/

var (
	Conf = &Config{}
)

func InitConfig() {
	confFile := "server.yaml"
	initConfig(Conf, filepath.Join(common.RootPath, "conf", confFile), "yaml")
}

// initConfig 将配置文件映射城结构体
func initConfig(c any, configPath, confType string) {
	// 判断文件是否存在
	if !fileutil.IsExist(configPath) {
		panic(fmt.Errorf("配置文件不存在: %s \n", configPath))
	}

	v := viper.New()
	v.SetConfigFile(configPath)
	v.SetConfigType(confType)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("无效的配置文件: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置文件变更:", e.Name)
		if err = v.Unmarshal(c); err != nil {
			panic(err)
		}
	})
	if err = v.Unmarshal(c); err != nil {
		panic(err)
	}
	log.Println("配置加载成功...", configPath)
}

type Config struct {
	Server    Server    `yaml:"server"`
	DB        DB        `yaml:"DB"`
	Redis     Redis     `yaml:"Redis"`
	Logs      Logs      `yaml:"Logs"`
	Jwt       Jwt       `yaml:"Jwt"`
	Websocket Websocket `yaml:"Websocket"`
	Mail      Mail      `yaml:"mail"`
}

type Server struct {
	Debug        bool   `yaml:"debug"`
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	ReadTimeout  int    `yaml:"ReadTimeout"`
	WriteTimeout int    `yaml:"WriteTimeout"`
	IdleTimeout  int    `yaml:"IdleTimeout"`
}

type DB struct {
	Enable        bool   `yaml:"Enable"`
	Type          string `yaml:"Type"`
	DSN           string `yaml:"DSN"`
	DbLog         string `yaml:"DbLog"`
	MaxLifetime   int    `yaml:"MaxLifetime"`
	MaxIdletime   int    `yaml:"MaxIdletime"`
	MaxOpenConns  int    `yaml:"MaxOpenConns"`
	MaxIdleConns  int    `yaml:"MaxIdleConns"`
	SlowThreshold int    `yaml:"SlowThreshold"`
}

type Logs struct {
	Level string `yaml:"Level"`
	Path  string `yaml:"Path"`
}

type Redis struct {
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	Password     string `yaml:"password"`
	DB           int    `yaml:"db"`
	MaxIdle      int    `yaml:"maxIdle"`
	MaxActive    int    `yaml:"maxActive"`
	IdleTimeout  int    `yaml:"idleTimeout"`
	PoolSize     int    `yaml:"poolSize"`
	MinIdleConns int    `yaml:"minIdleConns"`
}

type Jwt struct {
	JwtTokenSignKey string `yaml:"JwtTokenSignKey"`
	ExpiresTime     string `json:"ExpiresTime" yaml:"ExpiresTime"` // 过期时间
}

type Websocket struct {
	ReadDeadline          int `yaml:"ReadDeadline"`
	WriteDeadline         int `yaml:"WriteDeadline"`
	Start                 int `yaml:"Start"`
	WriteReadBufferSize   int `yaml:"WriteReadBufferSize"`
	MaxMessageSize        int `yaml:"MaxMessageSize"`
	PingPeriod            int `yaml:"PingPeriod"`
	HeartbeatFailMaxTimes int `yaml:"HeartbeatFailMaxTimes"`
}
type Mail struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
