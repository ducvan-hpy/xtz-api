package tzkt

import (
	"context"

	"github.com/ducvan-hpy/xtz-api/internal/domain/models"
)

type TzktSDK interface {
	GetDelegations(ctx context.Context) ([]models.Delegation, error)
}
