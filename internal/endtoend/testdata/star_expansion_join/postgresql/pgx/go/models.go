// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package querytest

import (
	"database/sql"
)

type Bar struct {
	C sql.NullString
	D sql.NullString
}

type Foo struct {
	A sql.NullString
	B sql.NullString
}
