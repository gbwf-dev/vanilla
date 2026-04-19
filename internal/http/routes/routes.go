package routes

import (
	"gravel/internal/http/handlers"
	"gravel/internal/services/vite"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
)

func Routes(router *fiber.App) {
	router.Get("/api/health", handlers.Health)
	router.Use("/api/", func(c fiber.Ctx) error { return c.SendStatus(fiber.StatusNotFound) })
	router.Use(static.New("", static.Config{FS: vite.FS}))
	router.Use(func(c fiber.Ctx) error { return c.SendFile("/index.html", fiber.SendFile{FS: vite.FS}) })
}
