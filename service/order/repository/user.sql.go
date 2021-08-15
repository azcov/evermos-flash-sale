package repository

import (
	"context"
	"database/sql"

	"github.com/azcov/evermos-flash-sale/constant"
	"github.com/azcov/evermos-flash-sale/domain/entity"
)

const getuserByEmail = `-- name: GetuserByEmail :one
SELECT id, name
FROM users
WHERE email = $1
`

func (q *Queries) GetuserByEmail(ctx context.Context, email string) (*entity.GetuserByEmailRow, error) {
	row := q.queryRow(ctx, q.getuserByEmailStmt, getuserByEmail, email)
	var i entity.GetuserByEmailRow
	err := row.Scan(&i.ID, &i.Name)
	if err == sql.ErrNoRows {
		return nil, constant.ErrorUserNotFound
	}
	return &i, err
}
