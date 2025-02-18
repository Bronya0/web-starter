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
