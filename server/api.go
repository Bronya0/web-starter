package main

import (
	"fmt"
	"web-starter/server/handler"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	debug = true
)

func main() {
	e := echo.New()
	if debug {
		Logger.Debug("Starting server In Debug Mode...")
	} else {
		Logger.Info("Starting server In Production Mode...")
		NewDB()
	}

	// 中间件
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Use(echoLogger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.Secure())

	// ================ API ===================
	g := e.Group("")

	restApi := &handler.RestApi{}
	g.GET("/:id", restApi.Get)
	g.POST("/post", restApi.Post)

	// 启动服务器
	fmt.Println("Starting server... http://localhost:8080")
	e.Logger.Fatal(e.Start(":8080")) // 监听8080端口
}
