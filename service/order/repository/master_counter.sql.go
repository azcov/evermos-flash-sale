package repository

import (
	"context"
	"database/sql"

	"github.com/azcov/evermos-flash-sale/constant"
	"github.com/azcov/evermos-flash-sale/domain/entity"
)

const getCounterByCounterID = `-- name: GetCounterByCounterID :one
SELECT counter, prefix, updated_at
FROM master_counters
WHERE id = $1
`

func (q *Queries) GetCounterByCounterID(ctx context.Context, counterID int32) (*entity.GetCounterByCounterIDRow, error) {
	row := q.queryRow(ctx, q.getCounterByCounterIDStmt, getCounterByCounterID, counterID)
	var i entity.GetCounterByCounterIDRow
	err := row.Scan(&i.Counter, &i.Prefix, &i.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, constant.ErrorCounterNotFound
	}
	return &i, err
}

const updateCounterByCounterID = `-- name: UpdateCounterByCounterID :exec
UPDATE master_counters
SET counter = $1,
    updated_at = $2
WHERE id = $3
`

func (q *Queries) UpdateCounterByCounterID(ctx context.Context, arg *entity.UpdateCounterByCounterIDParams) error {
	_, err := q.exec(ctx, q.updateCounterByCounterIDStmt, updateCounterByCounterID, arg.Counter, arg.UpdatedAt, arg.CounterID)
	return err
}
