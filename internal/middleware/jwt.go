package middleware

import (
	"github.com/labstack/echo/v4"
	"strings"
	"web-starter/internal/model/resp"
	"web-starter/internal/service"
)

func JWTAuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return resp.ErrorAuth(c)
			}
			// 按空格分割
			parts := strings.SplitN(authHeader, " ", 2)
			if !(len(parts) == 2 && parts[0] == "Bearer") {
				return resp.ErrorAuth(c)
			}
			// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它；也会自动校验过期时间
			payload, err := service.ParseToken(parts[1])
			if err != nil {
				return resp.ErrorAuth(c)
			}
			// 将当前请求的username信息保存到请求的上下文c上
			c.Set("UserID", payload.UserID)
			return next(c)
		}
	}
}
