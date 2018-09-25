package crons

import (
	"fmt"

	"github.com/andyvauliln/prediction-markets-api/app/app"
	"github.com/robfig/cron"
)

// tickStreamingCron takes instance of cron.Cron and adds tickStreaming
// crons according to the durations mentioned in config/app.yaml file
func (s *CronService) tickStreamingCron(c *cron.Cron) {
	for unit, durations := range app.Config.TickDuration {
		for _, duration := range durations {
			schedule := getCronScheduleString(unit, duration)
			c.AddFunc(schedule, s.tickStream(unit, duration))
		}
	}
}

// tickStream function fetches latest tick based on unit and duration for each Market
// and broadcasts the tick to the client subscribed to Market's respective channel
func (s *CronService) tickStream(unit string, duration int64) func() {
	return func() {

	}
}

// getCronScheduleString converts unit and duration to schedule string used for
// cron addFunc to schedule crons
func getCronScheduleString(unit string, duration int64) string {
	switch unit {

	case "sec":
		return fmt.Sprintf("*/%d * * * * *", duration)

	case "min":
		return fmt.Sprintf("0 */%d * * * *", duration)

	case "hour":
		return fmt.Sprintf("0 0 %d * * *", duration)

	case "day":
		return fmt.Sprintf("@daily")

	case "week":
		return fmt.Sprintf("0 0 0 * * */%d", duration)

	case "month":
		return fmt.Sprintf("0 0 0 */%d * *", duration)

	case "year":
		return fmt.Sprintf("@yearly")

	default:
		panic(fmt.Errorf("Invalid unit please try again"))
	}
}
