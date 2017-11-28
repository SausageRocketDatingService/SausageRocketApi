package main

import (
	"log"
	"os"

	"github.com/SausageRocketDatingService/SausageRocketApi/config"
	"github.com/SausageRocketDatingService/SausageRocketApi/handlers"
	"github.com/joho/godotenv"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())

	orm, dbError := config.InitDB()
	if dbError != nil {
		app.Logger().Fatalf("Failed to initialize DB: %v", dbError)
	}

	dbError = config.CreateDBTables(orm)

	if dbError != nil {
		app.Logger().Fatalf("Failed to initialize DB Tables: %v", dbError)
	}

	iris.RegisterOnInterrupt(func() {
		orm.Close()
	})

	app.Get("/", handlers.HelloWorld)

	app.Run(iris.Addr(":" + port))
}
