package handlers

import (
	"encoding/json"
	"fmt"
	"slot-golang/slots"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func BetHandler(c *fiber.Ctx) error {
	amount, err := strconv.ParseFloat(c.Query("betAmount"), 32)
	if err != nil {
		return c.SendString(err.Error())
	}
	numberOfLines, err := strconv.Atoi(c.Query("lines"))
	if err != nil {
		return c.SendString(err.Error())
	}

	result := slots.SlotV1Bet(float32(amount), numberOfLines)

	data, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}

	return c.SendString(string(data))
}
