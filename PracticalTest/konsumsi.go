package main

import (
	"encoding/json"
	"log"

	"github.com/go-resty/resty/v2"
)

// Struct untuk master konsumsi
type MasterKonsumsi struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	MaxPrice int    `json:"maxPrice"`
}

// Fungsi untuk mengambil data konsumsi dari API eksternal
func GetMasterKonsumsi() ([]MasterKonsumsi, error) {
	client := resty.New()
	resp, err := client.R().Get("https://6686cb5583c983911b03a7f3.mockapi.io/api/dummy-data/masterJenisKonsumsi")
	if err != nil {
		log.Println("Error fetching master konsumsi:", err)
		return nil, err
	}

	var konsumsiList []MasterKonsumsi
	err = json.Unmarshal(resp.Body(), &konsumsiList)
	if err != nil {
		log.Println("Error unmarshaling konsumsi data:", err)
		return nil, err
	}

	return konsumsiList, nil
}
