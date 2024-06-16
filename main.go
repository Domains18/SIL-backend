package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)




func  main(){
	app := fiber.New()




	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatalf("failed to start server: %v", app.Listen(":3000"))
}