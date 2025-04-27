package middleware

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"web-starter/internal/model/resp"
	"web-starter/internal/utils/glog"
)

// CustomRecovery 自定义错误(panic等)拦截中间件、对可能发生的错误进行拦截、统一记录
func CustomRecovery() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				if err := recover(); err != nil {
					errStr := fmt.Sprintf("%v", err)
					glog.Log.Error("【CustomRecovery】主协程内部错误：", errors.New(errStr))
					resp.Error(c, "", errStr)
				}
			}()
			return next(c)
		}
	}
}
