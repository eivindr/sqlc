// Code generated by sqlc. DO NOT EDIT.
// source: hstore.sql

package hstore

import (
	"context"
)

const listBar = `-- name: ListBar :many
SELECT bar FROM foo
`

func (q *Queries) ListBar(ctx context.Context) ([]interface{}, error) {
	rows, err := q.db.QueryContext(ctx, listBar)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []interface{}
	for rows.Next() {
		var bar interface{}
		if err := rows.Scan(&bar); err != nil {
			return nil, err
		}
		items = append(items, bar)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listBaz = `-- name: ListBaz :many
SELECT baz FROM foo
`

func (q *Queries) ListBaz(ctx context.Context) ([]interface{}, error) {
	rows, err := q.db.QueryContext(ctx, listBaz)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []interface{}
	for rows.Next() {
		var baz interface{}
		if err := rows.Scan(&baz); err != nil {
			return nil, err
		}
		items = append(items, baz)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
