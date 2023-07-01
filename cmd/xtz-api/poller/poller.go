package poller

import (
	"context"
	"log"
	"time"

	"github.com/ducvan-hpy/xtz-api/internal/domain/repository"
)

type Poller struct {
	tzktAPI    string
	interval   time.Duration
	repository *repository.Repository
}

func New(api string, interval time.Duration, repo *repository.Repository) *Poller {
	return &Poller{
		tzktAPI:    api,
		interval:   interval,
		repository: repo,
	}
}

func (p *Poller) Start(ctx context.Context) {
	log.Printf("Setup TzKT API Poller with interval: %v", p.interval)
	ticker := time.NewTicker(p.interval)
	for range ticker.C {
		p.pollDelegations(ctx)
	}
}

func (p *Poller) pollDelegations(ctx context.Context) {
	log.Println("Polling TzKT API:", p.tzktAPI)
	// GetDelegations
}
