// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package querytest

import (
	"context"
)

const plusPositionalCast = `-- name: PlusPositionalCast :one
SELECT plus($1, $2::INTEGER)
`

type PlusPositionalCastParams struct {
	A       int32
	Column2 int32
}

func (q *Queries) PlusPositionalCast(ctx context.Context, arg PlusPositionalCastParams) (int32, error) {
	row := q.db.QueryRow(ctx, plusPositionalCast, arg.A, arg.Column2)
	var plus int32
	err := row.Scan(&plus)
	return plus, err
}
