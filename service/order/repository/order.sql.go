package repository

import (
	"context"

	"github.com/azcov/evermos-flash-sale/domain/entity"
)

const createOrder = `-- name: CreateOrder :exec
INSERT INTO orders (order_no, user_id, total_price, created_at)
VALUES ($1, $2, $3, $4)
RETURNING id
`

func (q *Queries) CreateOrder(ctx context.Context, arg *entity.CreateOrderParams) (int32, error) {
	row := q.queryRow(ctx, q.createOrderStmt, createOrder,
		arg.OrderNo,
		arg.UserID,
		arg.TotalPrice,
		arg.CreatedAt,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}
