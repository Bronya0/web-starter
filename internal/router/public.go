package router

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
	"web-starter/internal/api"
	"web-starter/internal/api/v1/auth"
)

// addPublicRouter  公开的路由
func addPublicRouter(e *echo.Echo) *echo.Echo {
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"love you": time.Now().Format(time.DateTime),
		})
	})

	publicApi := e.Group("/api/public")
	{
		publicApi.GET("/father", api.TestGorm)
		publicApi.POST("/jwt-login", auth.Login)
	}

	return e
}
