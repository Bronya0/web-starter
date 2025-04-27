package middleware

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"time"
	"web-starter/internal/utils/glog"
)

func SlowTimeMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			startTime := time.Now()
			err := next(c)
			endTime := time.Now()
			latency := endTime.Sub(startTime)

			if latency.Seconds() > 1 { // 设置阈值，超过1秒则认为是慢接口
				glog.Log.Warn(fmt.Sprintf("【SlowTimeMiddleware】%v %v %v", c.Request().Method, c.Request().URL, latency))
			}
			return err
		}
	}
}
