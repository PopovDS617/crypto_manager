package token

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	"ms_price_saver/internal/client/db"
	"ms_price_saver/internal/model"
	"ms_price_saver/internal/repository"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.TokenRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, data *model.MessageQueueTokenData) (int64, error) {

	var (
		tableName       = data.TokenData.Symbol
		priceColumn     = "price"
		createdAtColumn = "created_at"
	)

	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(priceColumn, createdAtColumn).
		Values(data.TokenData.Price, data.Timestamp).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     fmt.Sprintf("%v_repository.Create", data.TokenData.Symbol),
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
