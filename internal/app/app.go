package main

import (
	"fmt"
	"math/rand/v2"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/alivecheck", func(c *fiber.Ctx) error {
		return c.SendString("ok")
	})

	app.Post("/order", func(c *fiber.Ctx) error {
		price := rand.IntN(1000)
		fmt.Printf("[Order] Received order, total price: %d\n", price)
		response_json := fmt.Sprintf("{\"total_price\": %d}", price)
		return c.SendString(response_json)
	})

	app.Get("/payment_status", func(c *fiber.Ctx) error {
		return c.SendString("{\"payment_valid\": true}")
	})

	app.Put("/payment", func(c *fiber.Ctx) error {
		rand_success := rand.IntN(5)
		success := rand_success != 0
		if success {
			fmt.Print("[Payment] Payment is successful!\n")
			return c.SendString("{\"payment_status\": true}")
		}
		fmt.Print("[Payment] Payment failed!\n")
		return c.SendString("{\"payment_status\": false}")
	})

	app.Listen(":6969")
}
