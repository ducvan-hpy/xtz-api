package persistence

import (
	"context"
	"log"

	"github.com/ducvan-hpy/xtz-api/internal/domain/model"
)

type InMemoryStorage struct {
	delegationsByYear map[int][]model.Delegation
	minYear           int
	maxYear           int
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		delegationsByYear: make(map[int][]model.Delegation),
		minYear:           10000,
		maxYear:           0,
	}
}

func (ims *InMemoryStorage) List(ctx context.Context, year *int) []model.Delegation {
	if year != nil {
		log.Printf("Filter on year %d", *year)
		delegations, ok := ims.delegationsByYear[*year]
		if !ok {
			return []model.Delegation{}
		}
		return delegations
	}

	delegations := make([]model.Delegation, 0, 100000)
	for i := ims.minYear; i <= ims.maxYear; i++ {
		delegationsBySelectedYear, ok := ims.delegationsByYear[i]
		if ok {
			delegations = append(delegations, delegationsBySelectedYear...)
		}
	}

	return delegations
}

func (ims *InMemoryStorage) Save(ctx context.Context, delegationsByYearToSave model.DelegationsByYearToSave) (int, int) {
	added := 0
	for year, delegations := range delegationsByYearToSave {
		ims.delegationsByYear[year] = append(ims.delegationsByYear[year], delegations...)
		if year < ims.minYear {
			ims.minYear = year
		}
		if year > ims.maxYear {
			ims.maxYear = year
		}
		added += len(delegations)
	}

	return ims.delegationsByYear[ims.maxYear][len(ims.delegationsByYear[ims.maxYear])-1].ID, added
}
