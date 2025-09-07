package models

import "database/sql"

type Author struct {
	Id      int            `db:"author_id"`
	Name    string         `db:"author_name"`
	Desc    sql.NullString `db:"author_desc"`
	Birth   sql.NullInt16  `db:"author_birth"`
	Country sql.NullString `db:"author_country"`
}