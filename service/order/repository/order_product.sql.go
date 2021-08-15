package repository

import (
	"context"

	"github.com/azcov/evermos-flash-sale/domain/entity"
	"github.com/lib/pq"
)

const createOrderProducts = `-- name: CreateOrderProducts :exec
INSERT INTO 
    order_products (order_id, product_id, price, qty, created_at)
SELECT
    $1, 
    unnest($2::int[]) AS product_id, 
    unnest($3::bigint[]) AS price, 
    unnest($4::int[]) AS qty, 
    $5
`

func (q *Queries) CreateOrderProducts(ctx context.Context, arg *entity.CreateOrderProductsParams) error {
	_, err := q.exec(ctx, q.createOrderProductsStmt, createOrderProducts,
		arg.OrderID,
		pq.Array(arg.ProductIds),
		pq.Array(arg.Prices),
		pq.Array(arg.Qtys),
		arg.CreatedAt,
	)
	return err
}
