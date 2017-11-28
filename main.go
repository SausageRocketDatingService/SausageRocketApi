package main

import (
	"github.com/SausageRocketDatingService/SausageRocketApi/handlers"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())

	app.Get("/", handlers.HelloWorld)

	app.Run(iris.Addr(":5000"))
}