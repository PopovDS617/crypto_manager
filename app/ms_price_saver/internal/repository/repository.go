package repository

import (
	"context"

	"ms_price_saver/internal/model"
)

type TokenRepository interface {
	Create(ctx context.Context, data *model.RepoTokenData) (int64, error)
}
