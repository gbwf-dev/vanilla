package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"gravel/internal/env"
	"gravel/internal/http/routes"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

type Validator func(any) error

func (fn Validator) Validate(out any) error { return fn(out) }

func main() {
	app := fiber.New(fiber.Config{
		AppName:  env.Get("APP_NAME"),
		Services: []fiber.Service{
			// Insert services here
		},
		StructValidator: Validator(validator.New().Struct),
	})
	app.Use(logger.New())
	routes.Routes(app)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := app.Listen(env.Get("LISTEN_ADDR"), fiber.ListenConfig{GracefulContext: ctx}); err != nil {
			log.Println(err)
		}
	}()

	<-ctx.Done()

	if err := app.Shutdown(); err != nil {
		log.Fatalln(err)
	}
}
