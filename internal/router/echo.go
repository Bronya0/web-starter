package router

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	echoMW "github.com/labstack/echo/v4/middleware"
	"net/http"
	"time"
	"web-starter/internal/config"
	"web-starter/internal/model/resp"
	"web-starter/internal/utils/glog"
)

// InitServer 加载配置文件的端口，启动echo服务，同时初始化路由
func InitServer() {
	e := NewEcho()

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

// Create 注册通用的路由
func NewEcho() *echo.Echo {
	e := New()
	// 放中间件前的路由,无需认证
	addPublicRouter(e)
	// 认证中间件
	addMiddleware(e)
	// 业务路由...需要认证Create
	addAuthRouter(e)
	return e
}

// New 初始化echo，这里放的中间件不能带认证
func New() *echo.Echo {
	e := echo.New()
	e.HideBanner = true

	e.Use(
		echoLogger(),
		echoRecover(),
		echoMW.RequestID(),
		echoMW.CSRF(),
		echoMW.Secure(),
	)
	e.HTTPErrorHandler = echoErrorHandler

	// 根据配置文件的debug初始化echo路由
	if config.Conf.Server.Debug == false {
		glog.Log.Info("【生产模式】")
	} else {
		glog.Log.Info("【调试模式】")
	}
	return e
}

// addMiddleware 这里放认证中间件
func addMiddleware(e *echo.Echo) {
	e.Use(
	//mw.ErrorLogger(),
	//mw.SlowTimeMiddleware(),
	)
}

// echoLogger echo日志记录
func echoLogger() echo.MiddlewareFunc {
	return echoMW.RequestLoggerWithConfig(echoMW.RequestLoggerConfig{
		// 设置需要记录的字段
		LogMethod:       true,
		LogURI:          true,
		LogStatus:       true,
		LogError:        true,
		LogProtocol:     true,
		LogLatency:      true, // 接口耗时
		LogResponseSize: true,
		//LogRequestID:    true,
		HandleError: true, // 转发错误到全局错误处理器
		LogValuesFunc: func(c echo.Context, v echoMW.RequestLoggerValues) error {
			if v.Error != nil {
				glog.Log.Error(fmt.Sprintf(`| %v | %v | %v "%v" | %v`, v.Status, v.Latency, v.Method, v.URI, v.Error.Error()))
			} else {
				glog.Log.Info(fmt.Sprintf(`| %v | %v | %v "%v"`, v.Status, v.Latency, v.Method, v.URI))

			}
			return nil
		},
	})
}

// echoRecover 主协程panic处理。
func echoRecover() echo.MiddlewareFunc {
	return echoMW.RecoverWithConfig(echoMW.RecoverConfig{
		StackSize: 2 << 10, // 2 KB
		LogErrorFunc: func(c echo.Context, err error, stack []byte) error {
			glog.Log.Errorf("内部服务错误: %v\n%s", err.Error(), string(stack))
			return resp.Error(c, "内部服务错误", nil)
		},
	})
}

// echoErrorHandler 统一错误处理
func echoErrorHandler(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}
	// Send response
	code := http.StatusInternalServerError
	var he *echo.HTTPError
	if errors.As(err, &he) {
		code = he.Code
	}

	glog.Log.Errorf("%+v", err)

	if c.Request().Method == http.MethodHead { // Issue #608
		err = c.NoContent(code)
	} else {
		err = c.JSON(code, map[string]any{
			"code": code,
			"msg":  "内部错误，详情见日志",
			"data": nil,
		})
	}
	if err != nil {
		glog.Log.Error(err.Error())
	}
}
