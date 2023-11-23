// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: query.sql

package querytest

import (
	"context"

	"github.com/jackc/pgconn"
)

const deleteBarByID = `-- name: DeleteBarByID :execresult
DELETE FROM bar WHERE id = $1
`

func (q *Queries) DeleteBarByID(ctx context.Context, id int32) (pgconn.CommandTag, error) {
	return q.db.Exec(ctx, deleteBarByID, id)
}
