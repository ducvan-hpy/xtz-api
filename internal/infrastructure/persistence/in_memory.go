package persistence

import (
	"context"

	"github.com/ducvan-hpy/xtz-api/internal/domain/models"
)

type InMemoryStorage struct {
	// TODO: Use map by year to speed filter up.
	delegations []models.Delegation
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{[]models.Delegation{}}
}

func (ims *InMemoryStorage) List(ctx context.Context) ([]models.Delegation, error) {
	ims.Save(ctx, models.Delegation{})
	return ims.delegations, nil
}

func (ims *InMemoryStorage) Save(ctx context.Context, delegation models.Delegation) error {
	ims.delegations = append(ims.delegations, delegation)
	return nil
}
