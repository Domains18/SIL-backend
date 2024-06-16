package main

import (
	"log"

	"github.com/Domains18/SIL-backend/database"
	"github.com/gofiber/fiber/v2"
)




func  main(){
	app := fiber.New()

	db, err := database.ConnectDatabase()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	err = database.MigrateDatabase(db)
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	app.Use(func (c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
		
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatalf("failed to start server: %v", app.Listen(":3000"))
}