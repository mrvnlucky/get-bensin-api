package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
)

type Fuel struct {
	ID    uint   `json:"id"`
	Name  string `json:"fuel_name"`
	Price string `json:"fuel_price"`
}

func main() {
	c := colly.NewCollector()
	var fuels []Fuel
	var lastID uint = 0
	var headers []string

	// Getting table headers
	c.OnHTML("table tbody tr:nth-child(1)", func(e *colly.HTMLElement) {
		e.ForEach("td, th", func(i int, el *colly.HTMLElement) {
			// if i == 0 {
			// 	return
			// }
			headers = append(headers, strings.Title(strings.ToLower(el.Text)))
		})
	})

	c.OnHTML("table tbody tr", func(e *colly.HTMLElement) {
		if e.Index == 0 {
			return
		}
		if !strings.Contains(e.Text, "Jakarta") && !strings.Contains(e.Text, "jakarta") {
			return
		}
		e.ForEach("td", func(i int, el *colly.HTMLElement) {
			if i == 0 {
				return
			}
			fuel := Fuel{
				ID:    lastID + 1,
				Name:  headers[i],
				Price: el.Text,
			}
			lastID++
			fuels = append(fuels, fuel)
		})
	})

	err := c.Visit("https://www.pertamina.com/id/news-room/announcement/daftar-harga-bahan-bakar-khusus-non-subsidi-tmt-1-februari-2024-zona-3")

	if err != nil {
		log.Fatal(err)
	}

	writeJSON(fuels)
}

func writeJSON(data []Fuel) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		fmt.Println("Unable to create JSON file")
		return
	}

	_ = os.WriteFile("fuels.json", file, 0644)
}
