package middleware

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"web-starter/internal/utils/glog"
)

func ErrorLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// 获取请求信息
			method := c.Request().Method
			url := c.Request().URL.String()
			statusCode := c.Response().Status

			// 记录日志
			if statusCode >= 500 {
				glog.Log.Error(fmt.Sprintf("【ErrorLogger】Method: %s, URL: %s, Status: %d", method, url, statusCode))
			}

			// 继续执行后续的中间件和路由
			return next(c)
		}
	}
}
