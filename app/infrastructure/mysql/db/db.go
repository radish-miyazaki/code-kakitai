package db

import (
	"context"
	"database/sql"
	"sync"
	"time"

	"github.com/radish-miyazaki/code-kakitai/infrastructure/mysql/db/db_gen"
)

const (
	maxRetries = 5
	delay      = 5 * time.Second

	QueriesKey = "queries"
)

var (
	once       sync.Once
	readOnce   sync.Once
	query      *db_gen.Queries
	readQuery  *db_gen.Queries
	dbConn     *sql.DB
	readDBConn *sql.DB
)

func GetQuery(ctx context.Context) *db_gen.Queries {
	q := getQueriesWithContext(ctx)
	if q != nil {
		return q
	}

	return query
}

func GetReadQuery(ctx context.Context) *db_gen.Queries {
	return readQuery
}

func GetDBConn() *sql.DB {
	return dbConn
}

func SetQuery(q *db_gen.Queries) {
	query = q
}

func SetDBConn(conn *sql.DB) {
	dbConn = conn
}

func WithQueries(ctx context.Context, queries *db_gen.Queries) context.Context {
	return context.WithValue(ctx, QueriesKey, queries)
}

func getQueriesWithContext(ctx context.Context) *db_gen.Queries {
	queries, ok := ctx.Value(QueriesKey).(*db_gen.Queries)
	if !ok {
		return nil
	}

	return queries
}
