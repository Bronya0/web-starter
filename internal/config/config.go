package config

import (
	"fmt"
	"log"
	"path/filepath"
	"web-starter/internal/common"

	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

/*
只允许依赖common包
*/

var (
	Conf = &Config{}
)

func InitConfig() {
	confFile := "server.toml"
	initConfig(Conf, filepath.Join(common.RootPath, "conf", confFile), "toml")
}

// initConfig 将配置文件映射城结构体
func initConfig(c any, configPath, confType string) {
	log.Println("初始化配置文件...", configPath)
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
	Server    Server    `toml:"Server"`
	DB        DB        `toml:"Db"`
	Redis     Redis     `toml:"Redis"`
	Log       Logs      `toml:"Log"`
	Jwt       Jwt       `toml:"Jwt"`
	Websocket Websocket `toml:"Websocket"`
	Mail      Mail      `toml:"Mail"`
}

type Server struct {
	Debug        bool   `toml:"Debug"`
	Host         string `toml:"Host"`
	Port         int    `toml:"Port"`
	ReadTimeout  int    `toml:"ReadTimeout"`
	WriteTimeout int    `toml:"WriteTimeout"`
	IdleTimeout  int    `toml:"IdleTimeout"`
}

type DB struct {
	Enable        bool   `toml:"Enable"`
	Type          string `toml:"Type"`
	DSN           string `toml:"DSN"`
	MaxLifetime   int    `toml:"MaxLifetime"`
	MaxIdletime   int    `toml:"MaxIdletime"`
	MaxOpenConns  int    `toml:"MaxOpenConns"`
	MaxIdleConns  int    `toml:"MaxIdleConns"`
	SlowThreshold int    `toml:"SlowThreshold"`
}

type Logs struct {
	Level string `toml:"Level"`
	Path  string `toml:"Path"`
	Db    string `toml:"Db"`
	Err   string `toml:"Err"`
}

type Redis struct {
	Host         string `toml:"Host"`
	Port         int    `toml:"Port"`
	Password     string `toml:"Password"`
	DB           int    `toml:"Db"`
	MaxIdle      int    `toml:"MaxIdle"`
	MaxActive    int    `toml:"MaxActive"`
	IdleTimeout  int    `toml:"IdleTimeout"`
	PoolSize     int    `toml:"PoolSize"`
	MinIdleConns int    `toml:"MinIdleConns"`
}

type Jwt struct {
	JwtTokenSignKey string `toml:"JwtTokenSignKey"`
	ExpiresTime     string `json:"ExpiresTime" toml:"ExpiresTime"` // 过期时间
}

type Websocket struct {
	ReadDeadline          int `toml:"ReadDeadline"`
	WriteDeadline         int `toml:"WriteDeadline"`
	Start                 int `toml:"Start"`
	WriteReadBufferSize   int `toml:"WriteReadBufferSize"`
	MaxMessageSize        int `toml:"MaxMessageSize"`
	PingPeriod            int `toml:"PingPeriod"`
	HeartbeatFailMaxTimes int `toml:"HeartbeatFailMaxTimes"`
}

type Mail struct {
	Host     string `toml:"Host"`
	Port     int    `toml:"Port"`
	Username string `toml:"Username"`
	Password string `toml:"Password"`
}
