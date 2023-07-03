package poller

import (
	"context"
	"log"
	"time"

	"github.com/ducvan-hpy/xtz-api/internal/domain/model"
	"github.com/ducvan-hpy/xtz-api/internal/domain/repository"
)

type DelegationsGetter func(context.Context) ([]model.Delegation, error)

type Poller struct {
	delegationGetter DelegationsGetter
	interval         time.Duration
	repository       *repository.Repository
}

func New(delegationGetter DelegationsGetter, interval time.Duration, repo *repository.Repository) *Poller {
	if delegationGetter == nil {
		log.Fatal("missing delegationGetter")
	}

	if repo == nil {
		log.Fatal("missing repository")
	}

	return &Poller{
		delegationGetter: delegationGetter,
		interval:         interval,
		repository:       repo,
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
	log.Println("Polling TzKT API")
	delegations, err := p.delegationGetter(ctx)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Got %d delegations", len(delegations))
	p.repository.Delegation.Save(ctx, delegations)
}
