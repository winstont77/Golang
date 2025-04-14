package controller

import(
	"GolangIrisMVC/service"
	"github.com/kataras/iris/v12"
	"io"
)

type UserController struct{
	userService service.UserService
}

func (userController *UserController) Get() string{
	return  userController.userService.GetUser()
}

func (c *UserController) Post(ctx iris.Context) string {
	body, err := io.ReadAll(ctx.Request().Body)
	if err != nil {
		return "Error reading body"
	}
	return c.userService.PostUser(string(body))
}