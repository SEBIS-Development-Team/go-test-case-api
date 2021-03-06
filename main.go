package main

import (
	"log"
	"test-case-api/database"
	"test-case-api/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Static("/public", "./public")
	database.ConnectDB()
	router.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
	//defer database.DB.Close()
}
