package controller

import (
	"GolangIrisMVC/model"
	"GolangIrisMVC/service"
	"github.com/kataras/iris/v12"
)

type UserController struct {
	userService service.UserService
}

func (userController *UserController) Get() string {
	return userController.userService.GetUser()
}

func (c *UserController) Post(ctx iris.Context) string {
	var user model.User
	if err := ctx.ReadJSON(&user); err != nil {
		return "Invalid JSON"
	}
	return c.userService.PostUser(user)
}