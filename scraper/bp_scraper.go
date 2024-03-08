package scraper

import (
	"fmt"
	"get-bensin/data"
	"get-bensin/util"
	"log"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

// ScrapeBP scrapes BP's website for fuel prices and stores them in the provided slice.
func ScrapeBP(fuels *[]data.Fuel) {
	names := []string{
		"BP Ultimate",
		"BP 92",
		"BP Diesel",
	}
	i := 0
	c := colly.NewCollector()

	// Find and parse the table
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:", r.URL.String())
	})
	c.OnHTML("table", func(e *colly.HTMLElement) {
		// Iterate over cells
		e.ForEach("tr:nth-child(n+2)", func(_ int, row *colly.HTMLElement) {
			price := strings.TrimSpace(row.ChildText("td:nth-child(2)"))
			fmt.Println(price)
			fuel := data.Fuel{
				Name:     names[i],
				Company:  "BP",
				Price:    util.ToIDR(price),
				DateTime: time.Now(),
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

	// Visit the URL
	c.Visit("https://www.bp.com/id_id/indonesia/home/produk-dan-layanan/spbu/harga.html")
}
