package main

import (
	"fmt"
	"strconv"
	"web-starter/server/handler"
	"web-starter/server/pkg"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	pkg.InitConfig("config.yaml")

	e := echo.New()
	e.HideBanner = true
	e.Validator = &pkg.CustomValidator{Validator: validator.New()}
	e.Use(pkg.EchoLogger())
	e.Use(middleware.Recover())

	// 连接db
	if pkg.Conf.Db.Enable {
		pkg.NewDB()
	}

	// debug模式
	if pkg.Conf.Debug {
		pkg.Logger.Info("Debug Mode!!")
		// 放开跨域
		e.Use(middleware.CORS())
	} else {
		pkg.Logger.Info("Production Mode...")
		e.Use(middleware.Secure())
	}

	// ================ API ===================
	g := e.Group("")

	restApi := &handler.RestApi{}
	g.GET("/:id", restApi.Get)
	g.POST("/post", restApi.Post)

	// 启动服务器
	port := strconv.Itoa(pkg.Conf.Serve.Port)
	fmt.Println("Starting server... http://localhost:" + port)
	e.Logger.Fatal(e.Start("127.0.0.1:" + port)) // 监听8080端口
}
