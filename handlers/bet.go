package handlers

import (
	"encoding/json"
	"fmt"
	"test-go/services"
	"test-go/slots"

	"github.com/gofiber/fiber/v2"
)

func BetHandler(c *fiber.Ctx) error {
	fmt.Println("bet_get_handler", c)
	amount := services.PlaceBet(5)
	fmt.Println(amount)
	result := slots.SlotV1(5)
	data, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}
	return c.SendString(string(data))
}
