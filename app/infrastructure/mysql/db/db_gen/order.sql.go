// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: order.sql

package db_gen

import (
	"context"
	"time"
)

const insertOrder = `-- name: InsertOrder :exec
INSERT INTO
    orders
    (
        id,
        user_id,
        total_amount,
        ordered_at
    )
VALUES
    (
        ?,
        ?,
        ?,
        ?
    )
`

type InsertOrderParams struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	TotalAmount int64     `json:"total_amount"`
	OrderedAt   time.Time `json:"ordered_at"`
}

func (q *Queries) InsertOrder(ctx context.Context, arg InsertOrderParams) error {
	_, err := q.db.ExecContext(ctx, insertOrder,
		arg.ID,
		arg.UserID,
		arg.TotalAmount,
		arg.OrderedAt,
	)
	return err
}

const orderFindByID = `-- name: OrderFindByID :one
SELECT
    id, user_id, total_amount, ordered_at, created_at, updated_at
FROM
    orders
WHERE
    id = ?
`

func (q *Queries) OrderFindByID(ctx context.Context, id string) (Orders, error) {
	row := q.db.QueryRowContext(ctx, orderFindByID, id)
	var i Orders
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.TotalAmount,
		&i.OrderedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
