package main

import (
	"context"
	"fmt"

	"github.com/driver005/gateway/registry"
	// "github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/cors"
	// "github.com/gofiber/fiber/v2/middleware/logger"
)

var ctx = context.Background()

func main() {
	r := registry.New(ctx)

	qrcode, p, err := r.Driver().PayNow()
	if err != nil {
		panic(err)
	}

	fmt.Println(p)
	fmt.Println(qrcode)
	// r.Setup()

	// public := fiber.New(fiber.Config{
	// 	ServerHeader:  "Fiber",
	// 	AppName:       "Test App v1.0.1",
	// 	WriteTimeout:  15 * time.Second,
	// 	ReadTimeout:   15 * time.Second,
	// 	StrictRouting: true,
	// })

	// r.RegisterRoutes(public)

	// public.Use(cors.New())
	// // public.Use(csrf.New())
	// public.Use(logger.New())
	// // public.Use(limiter.New())

	// public.Listen("localhost:80")
}
