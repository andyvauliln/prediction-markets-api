package crons

import (
	"github.com/andyvaulin/prediction-markets/services"
	"github.com/robfig/cron"
)

// CronService contains the services required to initialize crons
type CronService struct {
	marketService *services.MarketService
}

// NewCronService returns a new instance of CronService
func NewCronService(marketService *services.MarketService) *CronService {
	return &CronService{marketService}
}

// InitCrons is responsible for initializing all the crons in the system
func (s *CronService) InitCrons() {
	c := cron.New()
	s.tickStreamingCron(c)
	c.Start()
}
