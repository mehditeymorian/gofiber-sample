package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mehditeymorian/gofiber-sample/internal/cache"
	"github.com/mehditeymorian/gofiber-sample/internal/config"
	"github.com/mehditeymorian/gofiber-sample/internal/http/handler"
	"github.com/mehditeymorian/gofiber-sample/internal/http/middleware"
)

func main() {
	cfg := config.Load()

	cache := cache.New()

	app := fiber.New()
	app.Use(middleware.New)

	v1 := app.Group("/v1")
	handler.People{Cache: cache}.Register(v1)

	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatalf("error starting HTTP server: %v", err)
	}
}
