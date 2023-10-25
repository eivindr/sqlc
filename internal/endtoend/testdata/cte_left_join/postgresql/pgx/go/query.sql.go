// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: query.sql

package querytest

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const badQuery = `-- name: BadQuery :exec
WITH
	q
		AS (
			SELECT
				authors.name, authors.bio
			FROM
				authors
				LEFT JOIN fake ON authors.name = fake.name
		)
SELECT
	name, bio
FROM
	q AS c1
WHERE c1.name = $1
`

type BadQueryRow struct {
	Name string
	Bio  pgtype.Text
}

func (q *Queries) BadQuery(ctx context.Context, dollar_1 pgtype.Text) error {
	_, err := q.db.Exec(ctx, badQuery, dollar_1)
	return err
}
