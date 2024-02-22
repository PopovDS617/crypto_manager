package repository

import (
	"context"

	"ms_price_saver/internal/model"
)

type TokenRepository interface {
	Create(ctx context.Context, data *model.MessageQueueTokenData) (int64, error)
}
