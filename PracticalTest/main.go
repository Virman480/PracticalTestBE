package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/dashboard", func(c *fiber.Ctx) error {

		bookings, err := GetBookingList()
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Error fetching booking list"})
		}

		konsumsiList, err := GetMasterKonsumsi()
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Error fetching master konsumsi"})
		}

		dashboardData := []fiber.Map{}

		for _, booking := range bookings {
			totalCost := 0
			consumptionDetails := []fiber.Map{}

			for _, consumption := range booking.ListConsumption {
				var cost int
				for _, master := range konsumsiList {
					if master.Name == consumption.Name {
						cost = master.MaxPrice * booking.Participants
						totalCost += cost
						break
					}
				}

				consumptionDetails = append(consumptionDetails, fiber.Map{
					"name": consumption.Name,
					"cost": cost,
				})
			}

			dashboardData = append(dashboardData, fiber.Map{
				"roomName":           booking.RoomName,
				"officeName":         booking.OfficeName,
				"bookingDate":        booking.BookingDate,
				"participants":       booking.Participants,
				"consumptionDetails": consumptionDetails,
				"totalCost":          totalCost,
			})
		}

		return c.JSON(dashboardData)
	})

	log.Fatal(app.Listen(":8000"))
}
