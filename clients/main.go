package main

import (
	"flag"
	"fmt"
	"log"

	pb "root/proto"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	port = flag.Int("port", 3001, "The server port")
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	conn, err := grpc.Dial(fmt.Sprintf(":%d", *port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	client := pb.NewPostsServiceClient(conn)
	app.Get("/posts", func(c *fiber.Ctx) error {
		helo, err := client.HelloWorld(c.Context(), &pb.HelloWorldResponse{Message: "Hello World"})
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(helo)
	})

	log.Fatal(app.Listen(":3000"))
}