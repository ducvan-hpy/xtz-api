package poller

import (
	"context"
	"log"
	"time"

	"github.com/ducvan-hpy/xtz-api/internal/application/services/tzkt"
	"github.com/ducvan-hpy/xtz-api/internal/domain/repository"
)

type Poller struct {
	client          tzkt.TzktSDK
	limit           int
	interval        time.Duration
	intervalIdle    time.Duration
	intervalCurrent time.Duration
	repository      *repository.Repository
	lastID          int // Used for pagination
}

func New(client tzkt.TzktSDK, limit int, interval, intervalIdle time.Duration, repo *repository.Repository) *Poller {
	if client == nil {
		log.Fatal("missing TzKT delegations client")
	}

	if repo == nil {
		log.Fatal("missing repository")
	}

	return &Poller{
		client:          client,
		limit:           limit,
		interval:        interval,
		intervalIdle:    intervalIdle,
		intervalCurrent: interval,
		repository:      repo,
		lastID:          396673983447040,
	}
}

func (p *Poller) Start(ctx context.Context) {
	log.Printf("Setup TzKT API Poller with interval: %v", p.interval)
	ticker := time.NewTicker(p.interval)
	for range ticker.C {
		n := p.pollDelegations(ctx)
		if n == 0 {
			if p.intervalCurrent == p.interval {
				ticker.Reset(p.intervalIdle)
				p.intervalCurrent = p.intervalIdle
			}
		} else {
			if p.intervalCurrent == p.intervalIdle {
				ticker.Reset(p.interval)
				p.intervalCurrent = p.interval
			}
		}
		log.Printf("Got %d delegations, next poll in %v", n, p.intervalCurrent)
	}
}

func (p *Poller) pollDelegations(ctx context.Context) int {
	log.Println("Polling TzKT API")
	delegations, err := p.client.GetDelegations(ctx, p.limit, p.lastID)
	if err != nil {
		log.Println(err)
		return 0
	}
	p.lastID = p.repository.Delegation.Save(ctx, delegations)
	return len(delegations)
}
