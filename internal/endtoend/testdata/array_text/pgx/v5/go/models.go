// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0

package querytest

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Bar struct {
	Tags pgtype.Array[string]
}
