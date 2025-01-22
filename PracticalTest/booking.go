package main

import (
	"encoding/json"
	"log"

	"github.com/go-resty/resty/v2"
)

// Struct untuk data booking
type Booking struct {
	ID              string        `json:"id"`
	RoomName        string        `json:"roomName"`
	OfficeName      string        `json:"officeName"`
	BookingDate     string        `json:"bookingDate"`
	StartTime       string        `json:"startTime"`
	EndTime         string        `json:"endTime"`
	Participants    int           `json:"participants"`
	ListConsumption []Consumption `json:"listConsumption"`
}

// Struct untuk konsumsi dalam booking
type Consumption struct {
	Name string `json:"name"`
}

// Fungsi untuk mengambil data booking dari API eksternal
func GetBookingList() ([]Booking, error) {
	client := resty.New()
	resp, err := client.R().Get("https://66876cc30bc7155dc017a662.mockapi.io/api/dummy-data/bookingList")
	if err != nil {
		log.Println("Error fetching booking list:", err)
		return nil, err
	}

	var bookings []Booking
	err = json.Unmarshal(resp.Body(), &bookings)
	if err != nil {
		log.Println("Error unmarshaling booking data:", err)
		return nil, err
	}

	return bookings, nil
}
