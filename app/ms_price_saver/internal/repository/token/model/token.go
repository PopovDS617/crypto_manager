package model

import (
	"time"
)

type TokenPrice struct {
	ID        int64     `db:"id"`
	Price     float64   `db:"price"`
	CreatedAt time.Time `db:"created_at"`
}
