package routes

import (
	"flag"
	"fmt"
	"root/clients/database"
	pb "root/proto"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	port = flag.Int("port", 3001, "The server port")
	conn, _ = grpc.Dial(fmt.Sprintf(":%d", *port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	client = pb.NewPostsServiceClient(conn)
)

func HelloWorld(c *fiber.Ctx) (error) {
	helo, err := client.HelloWorld(c.Context(), &pb.HelloWorldResponse{Message: "Hello World"})
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(helo)
}

func GetAllPosts(c *fiber.Ctx) (error) {
	posts, err := client.GetAllPosts(c.Context(), &pb.Empty{})
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	data := posts.Posts
	database.DB.DB.Find(&data)
	return c.JSON(data)
}