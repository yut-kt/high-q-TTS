package entity

import "database/sql"

type TF struct {
	Line  int
	Str   string
	Class sql.NullString
}
