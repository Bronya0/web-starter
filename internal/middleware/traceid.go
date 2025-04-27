package middleware

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"web-starter/internal/common"
)

func TraceIDMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			traceID := c.Request().Header.Get("X-Request-ID") // 尝试从请求头获取 Trace ID
			if traceID == "" {
				traceID = generateTraceID() // 如果请求头中没有，则生成新的 Trace ID
			}

			c.Set(common.TraceIDKey, traceID) // 将 Trace ID 存储到 Echo Context 中

			// (可选) 将 Trace ID 添加到响应头
			c.Response().Header().Set("X-Request-ID", traceID)

			return next(c) // 继续处理请求
		}
	}
}

func generateTraceID() string {
	return uuid.New().String()
}
