//main.go

package main

import (
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Endpoint untuk dasAhboard
	app.Get("/dashboard", func(c *fiber.Ctx) error {
		// Ambil data booking
		bookings, err := GetBookingList()
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data booking"})
		}

		// Ambil data konsumsi
		konsumsiList, err := GetMasterKonsumsi()
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data konsumsi"})
		}

		// Filter berdasarkan tanggal (opsional)
		startDate := c.Query("startDate") // Format: YYYY-MM-DD
		endDate := c.Query("endDate")     // Format: YYYY-MM-DD
		if startDate != "" || endDate != "" {
			bookings = FilterBookingsByDate(bookings, startDate, endDate)
		}

		// Pagination
		page, _ := strconv.Atoi(c.Query("page", "1"))    // Default: page 1
		limit, _ := strconv.Atoi(c.Query("limit", "10")) // Default: 10 items per page
		bookings = PaginateBookings(bookings, page, limit)

		// Hasil akhir
		dashboardData := []fiber.Map{}

		// Proses setiap booking
		for _, booking := range bookings {
			totalCost := 0
			consumptionDetails := []fiber.Map{}

			// Proses konsumsi untuk setiap booking
			for _, consumption := range booking.ListConsumption {
				var cost int
				for _, master := range konsumsiList {
					if master.Name == consumption.Name {
						cost = master.MaxPrice * booking.Participants
						totalCost += cost
						break
					}
				}

				// Tambahkan detail konsumsi
				consumptionDetails = append(consumptionDetails, fiber.Map{
					"name": consumption.Name,
					"cost": cost,
				})
			}

			// Tambahkan data ruangan ke hasil akhir
			dashboardData = append(dashboardData, fiber.Map{
				"roomName":           booking.RoomName,
				"officeName":         booking.OfficeName,
				"bookingDate":        booking.BookingDate,
				"participants":       booking.Participants,
				"consumptionDetails": consumptionDetails,
				"totalCost":          totalCost,
			})
		}

		// Kembalikan hasil dalam bentuk JSON
		return c.JSON(dashboardData)
	})

	// Jalankan server di port 3000
	log.Fatal(app.Listen(":3000"))
}

// Fungsi untuk memfilter bookings berdasarkan tanggal
func FilterBookingsByDate(bookings []Booking, startDate string, endDate string) []Booking {
	filtered := []Booking{}

	// Parsing tanggal jika ada
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

	// Filter data
	for _, booking := range bookings {
		bookingDate, err := time.Parse(time.RFC3339, booking.BookingDate)
		if err != nil {
			continue
		}

		// Periksa apakah bookingDate berada dalam rentang startDate dan endDate
		if (start.IsZero() || !bookingDate.Before(start)) && (end.IsZero() || !bookingDate.After(end)) {
			filtered = append(filtered, booking)
		}
	}

	return filtered
}

// Fungsi untuk pagination bookings
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
