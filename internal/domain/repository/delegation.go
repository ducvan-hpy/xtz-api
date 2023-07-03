package repository

import (
	"context"

	"github.com/ducvan-hpy/xtz-api/internal/domain/model"
)

type Delegation interface {
	List(ctx context.Context) ([]model.Delegation, error)
	Save(ctx context.Context, delegations []model.Delegation) int
}
