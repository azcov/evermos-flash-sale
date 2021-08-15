package repository

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createOrderStmt, err = db.PrepareContext(ctx, createOrder); err != nil {
		return nil, fmt.Errorf("error preparing query CreateOrder: %w", err)
	}
	if q.createOrderProductsStmt, err = db.PrepareContext(ctx, createOrderProducts); err != nil {
		return nil, fmt.Errorf("error preparing query CreateOrderProducts: %w", err)
	}
	if q.getCounterByCounterIDStmt, err = db.PrepareContext(ctx, getCounterByCounterID); err != nil {
		return nil, fmt.Errorf("error preparing query GetCounterByCounterID: %w", err)
	}
	if q.getProductByProductIDForUpdateStmt, err = db.PrepareContext(ctx, getProductByProductIDForUpdate); err != nil {
		return nil, fmt.Errorf("error preparing query GetProductByProductIDForUpdate: %w", err)
	}
	if q.getProductsByProductIDForUpdateStmt, err = db.PrepareContext(ctx, getProductsByProductIDForUpdate); err != nil {
		return nil, fmt.Errorf("error preparing query GetProductsByProductIDForUpdate: %w", err)
	}
	if q.getuserByEmailStmt, err = db.PrepareContext(ctx, getuserByEmail); err != nil {
		return nil, fmt.Errorf("error preparing query GetuserByEmail: %w", err)
	}
	if q.updateCounterByCounterIDStmt, err = db.PrepareContext(ctx, updateCounterByCounterID); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateCounterByCounterID: %w", err)
	}
	if q.updateProductQtyStmt, err = db.PrepareContext(ctx, updateProductQty); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateProductQty: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createOrderStmt != nil {
		if cerr := q.createOrderStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createOrderStmt: %w", cerr)
		}
	}
	if q.createOrderProductsStmt != nil {
		if cerr := q.createOrderProductsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createOrderProductsStmt: %w", cerr)
		}
	}
	if q.getCounterByCounterIDStmt != nil {
		if cerr := q.getCounterByCounterIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getCounterByCounterIDStmt: %w", cerr)
		}
	}
	if q.getProductByProductIDForUpdateStmt != nil {
		if cerr := q.getProductByProductIDForUpdateStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getProductByProductIDForUpdateStmt: %w", cerr)
		}
	}
	if q.getProductsByProductIDForUpdateStmt != nil {
		if cerr := q.getProductsByProductIDForUpdateStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getProductsByProductIDForUpdateStmt: %w", cerr)
		}
	}
	if q.getuserByEmailStmt != nil {
		if cerr := q.getuserByEmailStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getuserByEmailStmt: %w", cerr)
		}
	}
	if q.updateCounterByCounterIDStmt != nil {
		if cerr := q.updateCounterByCounterIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateCounterByCounterIDStmt: %w", cerr)
		}
	}
	if q.updateProductQtyStmt != nil {
		if cerr := q.updateProductQtyStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateProductQtyStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                                  DBTX
	tx                                  *sql.Tx
	createOrderStmt                     *sql.Stmt
	createOrderProductsStmt             *sql.Stmt
	getCounterByCounterIDStmt           *sql.Stmt
	getProductByProductIDForUpdateStmt  *sql.Stmt
	getProductsByProductIDForUpdateStmt *sql.Stmt
	getuserByEmailStmt                  *sql.Stmt
	updateCounterByCounterIDStmt        *sql.Stmt
	updateProductQtyStmt                *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                                  tx,
		tx:                                  tx,
		createOrderStmt:                     q.createOrderStmt,
		createOrderProductsStmt:             q.createOrderProductsStmt,
		getCounterByCounterIDStmt:           q.getCounterByCounterIDStmt,
		getProductByProductIDForUpdateStmt:  q.getProductByProductIDForUpdateStmt,
		getProductsByProductIDForUpdateStmt: q.getProductsByProductIDForUpdateStmt,
		getuserByEmailStmt:                  q.getuserByEmailStmt,
		updateCounterByCounterIDStmt:        q.updateCounterByCounterIDStmt,
		updateProductQtyStmt:                q.updateProductQtyStmt,
	}
}
