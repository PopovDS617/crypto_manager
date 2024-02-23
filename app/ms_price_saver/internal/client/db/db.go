package db

import (
	"context"

	"github.com/jackc/pgx/v4"
)

// Client клиент для работы с БД
type Client interface {
	DB() DB
	Close() error
}

// Query обертка над запросом, хранящая имя запроса и сам запрос
// Имя запроса используется для логирования и потенциально может использоваться в других местах
type Query struct {
	Name     string
	QueryRaw string
}

// SQLExecer комбинирует NamedExecer и QueryExecer
type SQLExecer interface {
	QueryExecer
}

// QueryExecer интерфейс для работы с обычными запросами
type QueryExecer interface {
	QueryRowContext(ctx context.Context, q Query, args ...interface{}) pgx.Row
}

// Pinger интерфейс для проверки соединения с БД
type Pinger interface {
	Ping(ctx context.Context) error
}

// DB интерфейс для работы с БД
type DB interface {
	SQLExecer
	Pinger
	Close()
}
