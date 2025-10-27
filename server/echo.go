package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// API认证中间件
func apiAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		timestampStr := c.Request().Header.Get("X-Timestamp")
		clientToken := c.Request().Header.Get("X-Auth-Token")
		if timestampStr == "" || clientToken == "" {
			timestampStr = c.QueryParam("t")
			clientToken = c.QueryParam("token")
		}

		// 只允许：(AB有值且CD空) 或 (CD有值且AB空)
		if timestampStr == "" || clientToken == "" {
			Logger.Warn("认证头缺失", timestampStr, clientToken)
			return c.JSON(http.StatusForbidden, map[string]string{"error": "认证头缺失"})
		}

		timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
		if err != nil {
			Logger.Warn("无效的时间戳格式", timestampStr, timestamp)
			return c.JSON(http.StatusForbidden, map[string]string{"error": "无效的时间戳格式"})
		}

		// 检查时间戳是否在合理范围内（例如，前后60秒），防止重放攻击
		currentTimestamp := time.Now().Unix()
		if timestamp < currentTimestamp-1800 || timestamp > currentTimestamp+1800 {
			Logger.Warn("时间戳过期", timestamp, currentTimestamp)
			return c.JSON(http.StatusForbidden, map[string]string{"error": fmt.Sprintf("时间戳过期，你的本地时间与服务器时间差距过大，服务器时间：%v", currentTimestamp)})
		}

		// 在服务器端重新生成token
		h := sha256.New()
		h.Write([]byte(timestampStr + SharedSecret))
		serverToken := hex.EncodeToString(h.Sum(nil))

		// 比较token
		if clientToken != serverToken {
			Logger.Warn("无效的Token", clientToken, serverToken)
			return c.JSON(http.StatusForbidden, map[string]string{"error": "无效的Token"})
		}

		return next(c)
	}
}

func echoLogger() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		// 设置需要记录的字段
		LogMethod:       true,
		LogURI:          true,
		LogStatus:       true,
		LogError:        true,
		LogProtocol:     true,
		LogLatency:      true, // 接口耗时
		LogResponseSize: true,
		LogRemoteIP:     true,
		//LogRequestID:    true,
		HandleError: true, // 转发错误到全局错误处理器
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error != nil {
				Logger.Errorf(`| %v | %v | %v | %v "%v" | %v`, v.RemoteIP, v.Status, v.Latency, v.Method, v.URI, v.Error.Error())
			} else {
				Logger.Infof(`| %v | %v | %v | %v "%v"`, v.RemoteIP, v.Status, v.Latency, v.Method, v.URI)
			}
			return nil
		},
	})
}

// CustomValidator 是一个包装了 `validator.Validate` 的结构体
type CustomValidator struct {
	validator *validator.Validate
}

// Validate 是实现 echo.Validator 接口的方法
func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// 这里可以根据需要自定义错误信息的返回格式
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
