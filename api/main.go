package main

import (
	"fmt"
	"log"
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App){
	app.Get("/:url",routes.ResolveURL)
	app.Post("/api/v1",routes.ShortenURL)
}
func main(){
	err:=godotenv.Load()
	if err!=nil{
		fmt.Println(err)
	}
	app:=fiber.New()
	app.Use(logger.New())
	setupRoutes(app)
	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}

func ShortenURL(c *fiber.Ctx) error {
	body: = new(request)

	if err : =c.BodyParser(&body);err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"cannot parse JSON"})
	}
	//implement rate limiting

	//check if the input if an actual Url
	ig !govalidator.IsURL(body.URL){
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Invalid URL"})
	}

	//check for domain header
	if !helpers.RemoveDomainError(body.URL){
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{})
	}

	//enforce https ,ssl
	body.URL=helpers.EnforceHTTP(body.URL)
}