package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"web-starter/internal/model/resp"
)

//
//func Hello(c echo.Context) error {
//	time.Sleep(time.Second * 2)
//	return c.JSON(http.StatusOK, map[string]string{
//		"now": time.Now().Format(time.DateTime),
//	})
//}
//
//func HelloPost(c echo.Context) error {
//	type RegisterRequest struct {
//		Username string `json:"username" binding:"required"`
//		Email    string `json:"email" binding:"required,email"`
//		Age      uint8  `json:"age" binding:"gte=1,lte=120"`
//	}
//	var r RegisterRequest
//	if err := c.Bind(&r); err != nil {
//		return resp.Error(c, err.Error(), nil)
//	}
//	return c.JSON(http.StatusOK, map[string]string{
//		"title": "提交正确",
//	})
//}

type Page struct {
	Page int `form:"page" binding:"required,gte=1"`
	Size int `form:"size" binding:"required,gte=1"`
}

type FatherReq struct {
	FatherName string `json:"name"`
	FatherAge  int    `json:"age"`
	SonName    string `json:"son_name"`
	SonAge     int    `json:"son_age"`
}

func TestGorm(c echo.Context) error {
	var fathersWithSons []FatherReq
	var pager Page
	if err := c.Bind(&pager); err != nil {
		return resp.Error(c, err.Error(), nil)
	}
	return c.JSON(http.StatusOK, fathersWithSons)
}
