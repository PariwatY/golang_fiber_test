package main

import (
	"TestDevGolang/config"
	"TestDevGolang/handler"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db := config.InitialDB()
	_ = db
	app := fiber.New()

	app.Get("/:data", handler.GetPower)
	app.Listen(":3000")

}
