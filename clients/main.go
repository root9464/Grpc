package main

import (
	"log"
	"root/clients/database"
	"root/clients/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)


func main() {
	database.ConnectToDB()
	app := fiber.New()

	app.Use(logger.New())
	app.Get("/hello", routes.HelloWorld)
	app.Get("/posts", routes.GetAllPosts)
	log.Fatal(app.Listen(":3000"))
}