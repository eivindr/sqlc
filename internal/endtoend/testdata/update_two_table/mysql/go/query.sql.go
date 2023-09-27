// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: query.sql

package querytest

import (
	"context"
)

const deleteAuthor = `-- name: DeleteAuthor :exec
UPDATE
  authors,
  books
SET
  authors.deleted_at = now(),
  books.deleted_at = now()
WHERE
  books.is_amazing = 1
  AND authors.name = ?
`

func (q *Queries) DeleteAuthor(ctx context.Context, name string) error {
	_, err := q.db.ExecContext(ctx, deleteAuthor, name)
	return err
}
