// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: query.sql

package querytest

import (
	"context"
)

const setAuthor = `-- name: SetAuthor :exec
UPDATE  authors
SET     name = ?
WHERE   id = ?
`

type SetAuthorParams struct {
	Name string
	ID   int32
}

func (q *Queries) SetAuthor(ctx context.Context, arg SetAuthorParams) error {
	_, err := q.db.ExecContext(ctx, setAuthor, arg.Name, arg.ID)
	return err
}
