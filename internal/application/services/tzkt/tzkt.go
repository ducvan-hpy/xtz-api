package tzkt

import (
	"context"

	"github.com/ducvan-hpy/xtz-api/internal/domain/model"
)

type TzktSDK interface {
	GetDelegations(ctx context.Context, limit, lastID int) ([]model.Delegation, error)
}
