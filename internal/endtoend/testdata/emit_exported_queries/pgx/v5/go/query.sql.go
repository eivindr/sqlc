// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: query.sql

package querytest

import (
	"context"
)

const UpdateBarID = `-- name: UpdateBarID :exec
UPDATE bar SET id = $1 WHERE id = $2
`

type UpdateBarIDParams struct {
	ID   int32
	ID_2 int32
}

func (q *Queries) UpdateBarID(ctx context.Context, arg UpdateBarIDParams) error {
	_, err := q.db.Exec(ctx, UpdateBarID, arg.ID, arg.ID_2)
	return err
}
