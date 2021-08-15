package entity

import "database/sql"

type User struct {
	ID        int32          `json:"id"`
	Name      string         `json:"name"`
	Phone     sql.NullString `json:"phone"`
	Email     string         `json:"email"`
	CreatedAt int64          `json:"created_at"`
}

type GetuserByEmailRow struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}
