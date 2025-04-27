package auth

import (
	"github.com/labstack/echo/v4"
	"web-starter/internal/model/req"
	"web-starter/internal/model/resp"
	"web-starter/internal/service"
	"web-starter/internal/utils/glog"
)

func Login(c echo.Context) error {
	// 用户发送用户名和密码过来
	var login req.LoginReq
	if err := c.Bind(&login); err != nil {
		return resp.Error(c, "非法参数", nil)
	}
	// 验证用户名和密码
	user, err := service.CheckUser(login.Username, login.Password)
	if err != nil {
		return resp.Error(c, "用户名或密码错误", nil)
	}
	// 黑名单。从redis里检索set，存在则返回错误
	isBlacklisted, err := service.CheckBlacklist(login.Username)
	if err != nil {
		return resp.Error(c, "服务器内部错误", nil)
	}
	if isBlacklisted {
		return resp.Error(c, "账户已被禁用", nil)
	}
	// 生成JWT
	tokenString, err := service.GenToken(user.Id)
	if err != nil {
		glog.Log.Error("生成token失败", err)
		return nil
	}
	return resp.Success(c, "登录成功", map[string]string{"token": tokenString})
}
