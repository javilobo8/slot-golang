package handlers

import (
	"encoding/json"
	"fmt"
	"slot-golang/pkg/slots"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func BetHandler(c *fiber.Ctx) error {
	amount, err := strconv.ParseFloat(c.Query("betAmount"), 64)
	if err != nil {
		return c.SendString(err.Error())
	}
	numberOfLines, err := strconv.Atoi(c.Query("lines"))
	if err != nil {
		return c.SendString(err.Error())
	}

	result := slots.SlotV1Bet(amount, numberOfLines)

	data, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}

	return c.SendString(string(data))
}
