package pkg

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/Bronya0/go-utils/fileutil"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// https://ysboke.cn/tool/yaml2go

var (
	Conf = &Config{}
)

func InitConfig(confFile string) {
	initConfig(Conf, filepath.Join(WorkDir, confFile), "yaml")
}

// initConfig 将配置文件映射城结构体
func initConfig(c any, configPath, confType string) {
	// 判断文件是否存在
	if !fileutil.Exists(configPath) {
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
	Db struct {
		Enable bool `yaml:"enable"`
	} `yaml:"db"`
	Debug bool `yaml:"debug"`
	Serve struct {
		Port int `yaml:"port"`
	} `yaml:"serve"`
}
