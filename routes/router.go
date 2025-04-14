// routes/router.go
package routes

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"GolangIrisMVC/controller"
	"GolangIrisMVC/service"
)

func RegisterRoutes(app *iris.Application) {
	// Healthcheck 路由
	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})

	// User 路由
	user := app.Party("/user")
	mvc.Configure(user, func(app *mvc.Application) {
		app.Register(
			service.IUserService(&service.UserService{}),
		)
		app.Handle(new(controller.UserController))
	})
}
