package router

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"time"
	"web-starter/internal/config"
	mw "web-starter/internal/middleware"
	"web-starter/internal/utils/glog"
)

// InitServer 加载配置文件的端口，启动echo服务，同时初始化路由
func InitServer() {
	e := Create()

	cfg := config.Conf.Server
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	srv := &http.Server{
		Addr:         addr,
		Handler:      e,
		ReadTimeout:  time.Duration(cfg.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(cfg.IdleTimeout) * time.Second,
	}
	glog.Log.Info("欢迎主人！服务运行地址：http://", addr)
	glog.Log.Fatal(e.StartServer(srv))
}

// printRegisteredRoutes 打印注册的路由信息
func printRegisteredRoutes(e *echo.Echo) {
	// 遍历注册的路由
	for _, route := range e.Routes() {
		// 输出路由信息
		fmt.Printf("%s %s, ", route.Method, route.Path)
	}
}

// CreateEngine 注册通用的路由
func Create() *echo.Echo {
	e := New()
	// 放中间件前的路由,无需认证
	addPublicRouter(e)
	// 中间件
	addMiddleware(e)
	// 业务路由...需要认证
	addAuthRouter(e)
	return e
}

func New() *echo.Echo {
	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:   true,
		LogURI:      true,
		LogError:    true,
		HandleError: true, // forwards error to the global error handler, so it can decide appropriate status code
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error != nil {
				glog.Log.Error(v.Method, v.URI, v.URIPath, v.Status, v.Error.Error())
			} else {
				glog.Log.Info(v.Method, v.URI, v.URIPath, v.Status)
			}
			return nil
		},
	}))
	e.Use(middleware.Recover())

	// 根据配置文件的debug初始化echo路由
	if config.Conf.Server.Debug == false {
		glog.Log.Info("【生产模式】")
	} else {
		glog.Log.Info("【调试模式】")
	}
	return e
}

func addMiddleware(e *echo.Echo) {
	e.Use(
		mw.TraceIDMiddleware(),
		mw.ErrorLogger(),
		//mw.CustomRecovery(),
		mw.SlowTimeMiddleware(),
	)
}
