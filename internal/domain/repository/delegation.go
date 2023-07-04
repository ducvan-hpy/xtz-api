package repository

import (
	"context"

	"github.com/ducvan-hpy/xtz-api/internal/domain/model"
)

type Delegation interface {
	List(ctx context.Context, year *int) []model.Delegation
	Save(ctx context.Context, delegationsByYearToSave model.DelegationsByYearToSave) (int, int)
}
