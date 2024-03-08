package scraper

import (
	"fmt"
	"get-bensin/data"
	"get-bensin/util"

	"github.com/go-co-op/gocron/v2"
)

// ScheduleScraperJob schedules the scraper job to run every month at 00:00
func ScheduleScraperJob(fuels *[]data.Fuel) {
	// ScrapePertamina scrapes data from Pertamina
	ScrapePertamina(fuels)
	// ScrapeShell scrapes data from Shell
	ScrapeShell(fuels)
	// ScrapeBP scrapes data from BP
	ScrapeBP(fuels)

	// Write the fuels data to JSON
	util.WriteJSON(fuels)

	s, _ := gocron.NewScheduler()

	defer func() {
		_ = s.Shutdown()
	}()

	// Create a new job that runs every month on the 1st at 00:00
	_, _ = s.NewJob(
		gocron.MonthlyJob(
			1,
			gocron.NewDaysOfTheMonth(1),
			gocron.NewAtTimes(
				gocron.NewAtTime(0, 0, 0),
			),
		),
		gocron.NewTask(
			func() {
				// Run scraper functions
				ScrapePertamina(fuels)
				ScrapeShell(fuels)
				ScrapeBP(fuels)

				// Print the fuels data
				fmt.Println(fuels)

				// Write the fuels data to JSON
				util.WriteJSON(fuels)
			},
		),
	)
}
