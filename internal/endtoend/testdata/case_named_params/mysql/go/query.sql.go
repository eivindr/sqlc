// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: query.sql

package querytest

import (
	"context"
	"database/sql"
)

const listAuthors = `-- name: ListAuthors :one
SELECT  id, username, email, name, bio
FROM    authors
WHERE   email = CASE WHEN ? = '' then NULL else ? END
        OR username = CASE WHEN ? = '' then NULL else ? END 
LIMIT   1
`

type ListAuthorsParams struct {
	Email    sql.NullString
	Username sql.NullString
}

func (q *Queries) ListAuthors(ctx context.Context, arg ListAuthorsParams) (Author, error) {
	row := q.db.QueryRowContext(ctx, listAuthors,
		arg.Email,
		arg.Email,
		arg.Username,
		arg.Username,
	)
	var i Author
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Name,
		&i.Bio,
	)
	return i, err
}
