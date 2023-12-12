package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/nugrhrizki/gabut/cmd/web/routes"
	"github.com/nugrhrizki/gabut/internal"
)

func main() {
	internal.Setup()

	app := fiber.New(fiber.Config{
		Prefork: internal.State.Prefork,
	})

	defer app.Shutdown()

	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		TimeZone:   "Asia/Jakarta",
		TimeFormat: "02-Jan-2006 15:04:05",
	}))

	routes.Setup(app).
		Listen(fmt.Sprintf(":%d", internal.State.Port))
}
