package scraper

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"

	"get-bensin/data"
	"get-bensin/util"
)

// ScrapeShell scrapes Shell fuel prices from their website.
func ScrapeShell(fuels *[]data.Fuel) {
	// names is a list of fuel names.
	names := []string{
		"Shell Super",
		"Shell V-Power",
		"Shell V-Power Diesel",
		"Shell Diesel Extra",
		"Shell V-Power Nitro+",
	}
	// i is an index variable.
	i := 0

	// c is a colly collector.
	c := colly.NewCollector()

	// OnRequest is a callback function that is triggered for each request.
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:", r.URL.String())
	})

	// OnHTML is a callback function that is triggered when an HTML element is found.
	c.OnHTML("table tbody tr", func(e *colly.HTMLElement) {
		// Check if the text contains "jakarta".
		if !strings.Contains(strings.ToLower(e.Text), "jakarta") {
			return
		}
		// Iterate over the child elements of the "td:nth-child(n+2)" element.
		e.ForEach("td:nth-child(n+2)", func(_ int, el *colly.HTMLElement) {
			// Extract the price text.
			price := strings.TrimSpace(el.Text)
			// Print the price.
			fmt.Println(price)
			// Create a fuel struct.
			fuel := data.Fuel{
				Name:     names[i],
				Company:  "Shell",
				Price:    util.ToIDR(price),
				DateTime: time.Now(),
			}
			// Increment the index.
			i++
			// Append the fuel to the slice.
			*fuels = append(*fuels, fuel)
		})
	})

	// OnScraped is a callback function that is triggered when the scraping is complete.
	c.OnScraped(func(r *colly.Response) {
		fmt.Println(fuels)
		fmt.Println("Finished scraping:", r.Request.URL)
	})

	// OnError is a callback function that is triggered when an error occurs.
	c.OnError(func(r *colly.Response, err error) {
		log.Fatalln("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// Visit the Shell website.
	c.Visit("https://www.shell.co.id/in_id/pengendara-bermotor/bahan-bakar-shell/harga-bahan-bakar-shell.html")
}
