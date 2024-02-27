package main

import (
	"fmt"
	"get-bensin/scraper"
	"get-bensin/types"
	"get-bensin/util"
)

var (
	fuels []types.Fuel
)

func main() {
	scraper.ScrapePertamina(&fuels)
	scraper.ScrapeShell(&fuels)

	fmt.Println(&fuels)
	util.WriteJSON(&fuels)
}
