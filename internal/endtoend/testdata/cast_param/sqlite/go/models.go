// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package querytest

import (
	"database/sql"
)

type MyTable struct {
	Invalid sql.NullBool
	Foo     sql.NullString
}
