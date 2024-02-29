package main

import (
	"get-bensin/api"
	"get-bensin/data"
)

var (
	fuels []data.Fuel
)

func main() {
	r := api.SetupRouter()
	api.RegisterFuelRoutes(r)

	r.Run(":8080")
	// scraper.ScrapePertamina(&fuels)
	// scraper.ScrapeShell(&fuels)
	// scraper.ScrapeBP(&fuels)

	// fmt.Println(&fuels)
	// util.WriteJSON(&fuels)
}
