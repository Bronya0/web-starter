package common

import (
	"crypto/tls"
	"github.com/go-resty/resty/v2"
	"log"
	"net/http"
	"os"
	"time"
)

/*
不可依赖项目包
*/

var (
	RootPath   = getWorkDir()
	HttpClient = initHttpClient()
	Cache      = make(map[string]any)
)

// 这里定义的常量，一般是具有错误代码+错误说明组成，一般用于接口返回
const ()

func getWorkDir() string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return wd
}

func initHttpClient() *resty.Client {
	return resty.New().
		SetTransport(&http.Transport{
			MaxIdleConnsPerHost: 10,
		}).
		SetTimeout(60 * time.Second).
		SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
}
