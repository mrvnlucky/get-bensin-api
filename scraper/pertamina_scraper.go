package scraper

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly/v2"

	"get-bensin/data"
	"get-bensin/util"
)

func ScrapePertamina(fuels *[]data.Fuel) {
	names := []string{
		"Pertamax Turbo",
		"Pertamax Green",
		"Pertamax",
		"Pertalite",
		"Pertamina Dex",
		"Dexlite",
		"BioSolar Non-Subsidi",
		"BioSolar Subsidi",
	}
	i := 0
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:", r.URL.String())
	})
	c.OnHTML(".card-body", func(e *colly.HTMLElement) {
		e.ForEach(".d-flex.justify-content-between", func(_ int, el *colly.HTMLElement) {
			if !strings.Contains(strings.ToLower(el.Text), "jakarta") {
				return
			}
			price := strings.TrimSpace(el.ChildText("label:last-child"))
			fuel := data.Fuel{
				Name:    names[i],
				Company: "Pertamina",
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
	c.Visit("https://mypertamina.id/fuels-harga")
}
