package handler

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/muhrobby/go-redirect/database"
	"github.com/muhrobby/go-redirect/model"
	"github.com/muhrobby/go-redirect/utils"
)

func CreateShortLink(c *fiber.Ctx) error {

	type Sl struct {
		LongLink string `json:"long_link"`
		NameLink string `json:"name_link"`
	}

	var ShrtLnk Sl

	shortlink := new(model.Shortlink)
	err := c.BodyParser(&ShrtLnk)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "gagal ambil data",
		})
	}

	if ShrtLnk.NameLink == "" {

		shortlink.NameLink = utils.RandomString(5)

	} else {
		shortlink.NameLink = strings.Replace(ShrtLnk.NameLink, " ", "-", -1)
	}

	isHttps := strings.HasPrefix(ShrtLnk.LongLink, "https://")

	if !isHttps {
		return c.JSON(fiber.Map{
			"message": "tidak di awali dengan https://",
		})

	}
	shortlink.LongLink = ShrtLnk.LongLink

	db := database.Connect()

	err = db.Create(&shortlink).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "gagal simpan data",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "berhasil simpan data",
		"data":    ShrtLnk,
	})
}

func RedirectShortLink(c *fiber.Ctx) error {
	namelink := c.Params("id")

	var Sl model.Shortlink

	err := database.Connect().Find(&Sl, "name_link = ?", namelink).Error

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "failed to find",
		})
	}

	fmt.Println(Sl.LongLink)

	return c.Redirect(Sl.LongLink, fiber.StatusFound)

}
