// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: query.sql

package querytest

import (
	"context"
)

const placeholder = `-- name: Placeholder :exec
SELECT baz from foo
`

func (q *Queries) Placeholder(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, placeholder)
	return err
}
