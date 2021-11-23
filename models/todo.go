package models

import "database/sql"

type Todo struct {
	ID          int64
	Title       sql.NullString
	Content     sql.NullString
	Status      sql.NullInt16
	OwnerId     sql.NullString
	CreatedTime sql.NullString
	UpdateTime  sql.NullString
}
