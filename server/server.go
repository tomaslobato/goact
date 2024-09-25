package server

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func New() {
	app := fiber.New()

	//frontend handler, define your frontend routes
	frontRoutes := []string{
		"/",
		"/about",
	}

	HandleFrontRoutes(app, frontRoutes)

	//regular backend route
	app.Get("/hello", HandleHelloWorld)

	//use PORT env variable or default :3000
	godotenv.Load(".env.local")
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	app.Listen(":" + port)
}
