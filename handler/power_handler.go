package handler

import (
	"TestDevGolang/config"
	"TestDevGolang/model"
	"github.com/gofiber/fiber/v2"
)

func GetPower(c *fiber.Ctx) error {
	var power model.Power
	sumData := c.Params("data")

	//select  active_power and power_input which sum equal sumData
	config.InitialDB().Raw("SELECT * FROM powers p WHERE active_power + power_input = ?", sumData).Scan(&power)
	return c.Status(200).JSON(power)
}
