package middleware

import (
	"gin-starter/internal/common"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//检查请求头中是否已存在 Trace ID: 通常，Trace ID 会通过请求头 (X-Request-ID 或自定义的 header) 传递。
//如果上游服务已经生成了 Trace ID 并传递下来，你应该 复用 这个 Trace ID，而不是重新生成。
//如果请求头中没有 Trace ID，则生成一个新的 Trace ID。
//将 Trace ID 存储到 Gin 的 Context 中: Gin 的 Context 是请求级别的上下文，可以在整个请求处理链中共享数据。 方便在 Handler 函数中访问。
//(可选) 将 Trace ID 添加到响应头中: 为了方便下游服务继续追踪，可以将 Trace ID 添加到响应头中，例如 X-Request-ID。

func TraceIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := c.Request.Header.Get("X-Request-ID") // 尝试从请求头获取 Trace ID
		if traceID == "" {
			traceID = generateTraceID() // 如果请求头中没有，则生成新的 Trace ID
		}

		c.Set(common.TraceIDKey, traceID) // 将 Trace ID 存储到 Gin Context 中

		// (可选) 将 Trace ID 添加到响应头
		c.Writer.Header().Set("X-Request-ID", traceID)

		c.Next() // 继续处理请求
	}
}

func generateTraceID() string {
	return uuid.New().String()
}
