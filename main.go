// main.go
package main

import (
	"github.com/kataras/iris/v12"
	"GolangIrisMVC/routes"
)

func main() {
	app := iris.New()

	// 呼叫封裝好的註冊路由方法
	routes.RegisterRoutes(app)

	app.Listen(":8080", iris.WithLogLevel("debug"))
}
