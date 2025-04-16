// main.go
package main

import (
	"log"
	"github.com/kataras/iris/v12"
	"GolangIrisMVC/routes"
	"GolangIrisMVC/infrastructure"
)

func main() {
	// 初始化資料庫連線
	infrastructure.InitDB()

	// 建立 Iris App
	app := iris.New()

	// 呼叫封裝好的註冊路由方法
	routes.RegisterRoutes(app)

	// 啟動伺服器
	err := app.Listen(":8080", iris.WithLogLevel("debug"))
	if err != nil {
		log.Fatal(err)
	}
}
