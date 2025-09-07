package models

import "database/sql"

type Genre struct {
	Id   int            `db:"genre_id"`
	Name string         `db:"genre_name"`
	Desc sql.NullString `db:"genre_desc"`
}