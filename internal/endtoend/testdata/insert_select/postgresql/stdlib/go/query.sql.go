// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: query.sql

package querytest

import (
	"context"
)

const insertSelect = `-- name: InsertSelect :exec
INSERT INTO foo (name, meta)
SELECT name, $1
FROM bar WHERE ready = $2
`

type InsertSelectParams struct {
	Meta  string
	Ready bool
}

func (q *Queries) InsertSelect(ctx context.Context, arg InsertSelectParams) error {
	_, err := q.db.ExecContext(ctx, insertSelect, arg.Meta, arg.Ready)
	return err
}
