// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: query.sql

package querytest

import (
	"context"
)

const multiFrom = `-- name: MultiFrom :many
SELECT email FROM bar, foo WHERE login = ?
`

func (q *Queries) MultiFrom(ctx context.Context, login string) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, multiFrom, login)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var email string
		if err := rows.Scan(&email); err != nil {
			return nil, err
		}
		items = append(items, email)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
