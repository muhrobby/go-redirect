package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhrobby/go-redirect/handler"
)

func Router(app *fiber.App) {

	app.Get("/:id", handler.RedirectShortLink)
	app.Post("/create", handler.CreateShortLink)

}
