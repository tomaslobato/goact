package server

import "github.com/gofiber/fiber/v2"

func HandleHelloWorld(c *fiber.Ctx) error {
	return c.JSON("Hello world")
}

func HandleFrontRoutes(app *fiber.App, routes []string) {
	app.Static("/", "./dist", fiber.Static{
		CacheDuration: 200,
	})

	for _, r := range routes {
		app.Get(r, func(c *fiber.Ctx) error {
			return c.Render("public/index.html	", fiber.Map{})
		})
	}
}
