package main

import (
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	
	app.Get("/dashboard", func(c *fiber.Ctx) error {
		
		bookings, err := GetBookingList()
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data booking"})
		}

		
		konsumsiList, err := GetMasterKonsumsi()
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data konsumsi"})
		}

		startDate := c.Query("startDate") 
		endDate := c.Query("endDate")     
		if startDate != "" || endDate != "" {
			bookings = FilterBookingsByDate(bookings, startDate, endDate)
		}

		
		page, _ := strconv.Atoi(c.Query("page", "1"))    
		limit, _ := strconv.Atoi(c.Query("limit", "10")) 
		bookings = PaginateBookings(bookings, page, limit)

		
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

	
	log.Fatal(app.Listen(":3000"))
}


func FilterBookingsByDate(bookings []Booking, startDate string, endDate string) []Booking {
	filtered := []Booking{}

	
	var start, end time.Time
	var err error
	if startDate != "" {
		start, err = time.Parse("2006-01-02", startDate)
		if err != nil {
			log.Println("Format startDate salah, gunakan YYYY-MM-DD")
		}
	}
	if endDate != "" {
		end, err = time.Parse("2006-01-02", endDate)
		if err != nil {
			log.Println("Format endDate salah, gunakan YYYY-MM-DD")
		}
	}

	
	for _, booking := range bookings {
		bookingDate, err := time.Parse(time.RFC3339, booking.BookingDate)
		if err != nil {
			continue
		}

		
		if (start.IsZero() || !bookingDate.Before(start)) && (end.IsZero() || !bookingDate.After(end)) {
			filtered = append(filtered, booking)
		}
	}

	return filtered
}


func PaginateBookings(bookings []Booking, page int, limit int) []Booking {
	start := (page - 1) * limit
	if start > len(bookings) {
		return []Booking{}
	}

	end := start + limit
	if end > len(bookings) {
		end = len(bookings)
	}

	return bookings[start:end]
}
