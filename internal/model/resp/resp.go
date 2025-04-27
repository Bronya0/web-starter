package resp

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Success(c echo.Context, msg string, data any) error {
	return c.JSON(http.StatusOK, map[string]any{
		"code": 2000,
		"msg":  msg,
		"data": data,
	})
}

func Error(c echo.Context, msg string, data any) error {
	return c.JSON(http.StatusOK, map[string]any{
		"code": 5000,
		"msg":  msg,
		"data": data,
	})
}

func ErrorAuth(c echo.Context) error {
	return c.JSON(http.StatusUnauthorized, map[string]any{
		"code": 4001,
		"msg":  "认证不通过",
		"data": nil,
	})
}