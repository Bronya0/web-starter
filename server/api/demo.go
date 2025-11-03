package api

import (
	"github.com/labstack/echo/v4"
)

type RestApi struct {
}

func (d *RestApi) Get(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(200, map[string]string{"message": id})
}

func (d *RestApi) Post(c echo.Context) error {
	var params struct {
		Files []struct {
			Filename string `json:"filename" validate:"required,min=5,max=20"`
		} `json:"files"`
		Title string `json:"title" validate:"required,min=5,max=20"` // min/max 对字符串来说是字符数量，兼容中文
		Email string `json:"email"    validate:"required,email"`
	}
	if err := c.Bind(&params); err != nil {
		return c.JSON(400, map[string]string{"msg": "参数错误"})
	}
	if err := c.Validate(params); err != nil {
		return c.JSON(404, map[string]string{"msg": "参数错误"})
	}
	return c.JSON(200, map[string]any{"msg": params})
}
