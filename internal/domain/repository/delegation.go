package repository

import (
	"context"

	"github.com/ducvan-hpy/xtz-api/internal/domain/models"
)

type Delegation interface {
	List(ctx context.Context) ([]models.Delegation, error)
	Save(ctx context.Context, delegations []models.Delegation) error
}
