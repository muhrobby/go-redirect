package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhrobby/go-redirect/database"
	"github.com/muhrobby/go-redirect/migration"
	"github.com/muhrobby/go-redirect/router"
)

func main() {
	app := fiber.New()

	database.Connect()
	migration.AutoMigrate()

	router.Router(app)
	app.Listen(":3030")
}
