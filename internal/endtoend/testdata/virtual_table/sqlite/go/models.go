// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package querytest

import (
	"database/sql"
)

type Ft struct {
	B string
}

type Tbl struct {
	A int64
	B sql.NullString
	C sql.NullString
	D sql.NullString
	E sql.NullInt64
}

type TblFt struct {
	B string
	C string
}
