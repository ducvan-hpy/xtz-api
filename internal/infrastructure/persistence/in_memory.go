package persistence

import (
	"context"

	"github.com/ducvan-hpy/xtz-api/internal/domain/model"
)

type InMemoryStorage struct {
	// TODO: Use map by year to speed filter up.
	delegations []model.Delegation
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{[]model.Delegation{}}
}

func (ims *InMemoryStorage) List(ctx context.Context) ([]model.Delegation, error) {
	return ims.delegations, nil
}

func (ims *InMemoryStorage) Save(ctx context.Context, delegation []model.Delegation) error {
	ims.delegations = append(ims.delegations, delegation...)
	return nil
}
