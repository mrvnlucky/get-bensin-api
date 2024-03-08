package main

import (
	"get-bensin/api"
	"get-bensin/data"
	"get-bensin/scraper"
)

var (
	fuels []data.Fuel
)

func main() {
	scraper.ScheduleScraperJob(&fuels)
	r := api.SetupRouter()
	r.Run(":8080")

	// scraper.ScrapePertamina(&fuels)
	// scraper.ScrapeShell(&fuels)
	// scraper.ScrapeBP(&fuels)

	// fmt.Println(&fuels)
	// util.WriteJSON(&fuels)
}
