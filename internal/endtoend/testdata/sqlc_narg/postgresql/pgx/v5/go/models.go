// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package querytest

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Foo struct {
	Bar      string
	MaybeBar pgtype.Text
}
