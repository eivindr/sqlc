// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package querytest

import (
	"context"
	"database/sql"
)

const getFirst = `-- name: GetFirst :many
SELECT val FROM first_view
`

func (q *Queries) GetFirst(ctx context.Context) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getFirst)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var val string
		if err := rows.Scan(&val); err != nil {
			return nil, err
		}
		items = append(items, val)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSecond = `-- name: GetSecond :many
SELECT val, val2 FROM second_view WHERE val2 = ?
`

func (q *Queries) GetSecond(ctx context.Context, val2 sql.NullInt64) ([]SecondView, error) {
	rows, err := q.db.QueryContext(ctx, getSecond, val2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SecondView
	for rows.Next() {
		var i SecondView
		if err := rows.Scan(&i.Val, &i.Val2); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
