package router

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

// addAuthRouter 需要认证的路由
func addAuthRouter(r *echo.Echo) *echo.Echo {
	authApi := r.Group("/api/v1")
	{
		authApi.GET("/", func(c echo.Context) error {
			return c.JSON(http.StatusOK, map[string]string{
				"love you": time.Now().Format(time.DateTime),
			})
		})
	}

	return r
}
