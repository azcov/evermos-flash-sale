package repository

import (
	"context"

	"github.com/azcov/evermos-flash-sale/constant"
	"github.com/azcov/evermos-flash-sale/domain/entity"
	"github.com/lib/pq"
)

const getProductByProductIDForUpdate = `-- name: GetProductByProductIDForUpdate :one
SELECT name, price, qty 
FROM products
WHERE id = $1
FOR UPDATE
`

func (q *Queries) GetProductByProductIDForUpdate(ctx context.Context, productID int32) (*entity.GetProductByProductIDForUpdateRow, error) {
	row := q.queryRow(ctx, q.getProductByProductIDForUpdateStmt, getProductByProductIDForUpdate, productID)
	var i entity.GetProductByProductIDForUpdateRow
	err := row.Scan(&i.Name, &i.Price, &i.Qty)
	if err != nil {
		return nil, constant.ErrorProductNotFound
	}
	return &i, err
}

const getProductsByProductIDForUpdate = `-- name: GetProductsByProductIDForUpdate :many
SELECT id, name, price, qty 
FROM products
WHERE id = ANY($1::int[])
FOR UPDATE
`

func (q *Queries) GetProductsByProductIDForUpdate(ctx context.Context, productIds []int32) (map[int32]*entity.GetProductByProductIDForUpdateRow, error) {
	rows, err := q.query(ctx, q.getProductsByProductIDForUpdateStmt, getProductsByProductIDForUpdate, pq.Array(productIds))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := map[int32]*entity.GetProductByProductIDForUpdateRow{}
	for rows.Next() {
		var i entity.GetProductByProductIDForUpdateRow
		if err := rows.Scan(&i.ID, &i.Name, &i.Price, &i.Qty); err != nil {
			return nil, err
		}
		items[i.ID] = &i
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	if len(productIds) != len(items) {
		return nil, constant.ErrorProductsNotFound
	}
	return items, nil
}

const updateProductQty = `-- name: UpdateProductQty :exec
UPDATE products
SET qty = $1
WHERE id = $2
`

func (q *Queries) UpdateProductQty(ctx context.Context, arg *entity.UpdateProductQtyParams) error {
	_, err := q.exec(ctx, q.updateProductQtyStmt, updateProductQty, arg.Qty, arg.ProductID)
	return err
}
