package repository

import (
	"context"
	"database/sql"

	"github.com/azcov/evermos-flash-sale/domain/entity"
)

type Repository interface {
	CreateOrder(ctx context.Context, arg *entity.CreateOrderParams) (int32, error)
	CreateOrderProducts(ctx context.Context, arg *entity.CreateOrderProductsParams) error
	GetCounterByCounterID(ctx context.Context, counterID int32) (*entity.GetCounterByCounterIDRow, error)
	GetProductByProductIDForUpdate(ctx context.Context, productID int32) (*entity.GetProductByProductIDForUpdateRow, error)
	GetProductsByProductIDForUpdate(ctx context.Context, productIds []int32) (map[int32]*entity.GetProductByProductIDForUpdateRow, error)
	GetuserByEmail(ctx context.Context, email string) (*entity.GetuserByEmailRow, error)
	UpdateCounterByCounterID(ctx context.Context, arg *entity.UpdateCounterByCounterIDParams) error
	UpdateProductQty(ctx context.Context, arg *entity.UpdateProductQtyParams) error

	// Tx Function
	WithTx(tx *sql.Tx) *Queries
	BeginTx(ctx context.Context) (*sql.Tx, error)
	RollbackTx(tx *sql.Tx) error
	CommitTx(tx *sql.Tx) error
}
