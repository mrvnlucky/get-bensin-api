package scraper

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly/v2"

	"get-bensin/types"
	"get-bensin/util"
)

func ScrapeShell(fuels *[]types.Fuel) {
	names := []string{
		"Super",
		"V-Power",
		"V-Power Diesel",
		"Diesel Extra",
		"V-Power Nitro+",
	}
	i := 0

	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:", r.URL.String())
	})
	c.OnHTML("table tbody tr", func(e *colly.HTMLElement) {
		if !strings.Contains(strings.ToLower(e.Text), "jakarta") {
			return
		}
		e.ForEach("td:nth-child(n+2)", func(_ int, el *colly.HTMLElement) {
			price := el.Text
			fuel := types.Fuel{
				Name:    names[i],
				Company: "Shell",
				Price:   util.ToIDR(price),
			}
			i++
			*fuels = append(*fuels, fuel)
		})
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println(fuels)
		fmt.Println("Finished scraping:", r.Request.URL)
	})
	c.OnError(func(r *colly.Response, err error) {
		log.Fatalln("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})
	c.Visit("https://www.shell.co.id/in_id/pengendara-bermotor/bahan-bakar-shell/harga-bahan-bakar-shell.html")
}
