// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: query.sql

package querytest

import (
	"context"
	"database/sql"
)

const testInnerLeft = `-- name: TestInnerLeft :many
SELECT a.a, b.b, c.c
FROM a
INNER JOIN b ON b.a_id = a.id
LEFT JOIN c ON c.a_id = a.id
`

type TestInnerLeftRow struct {
	A string
	B string
	C sql.NullString
}

func (q *Queries) TestInnerLeft(ctx context.Context) ([]TestInnerLeftRow, error) {
	rows, err := q.db.QueryContext(ctx, testInnerLeft)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TestInnerLeftRow
	for rows.Next() {
		var i TestInnerLeftRow
		if err := rows.Scan(&i.A, &i.B, &i.C); err != nil {
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

const testInnerLeftInnerLeft = `-- name: TestInnerLeftInnerLeft :many
SELECT a.a, b.b, c.c, d.d, e.e
FROM a
INNER JOIN b ON b.a_id = a.id
LEFT JOIN c ON c.a_id = a.id
INNER JOIN d ON d.a_id = a.id
LEFT JOIN e ON e.a_id = a.id
`

type TestInnerLeftInnerLeftRow struct {
	A string
	B string
	C sql.NullString
	D string
	E sql.NullString
}

func (q *Queries) TestInnerLeftInnerLeft(ctx context.Context) ([]TestInnerLeftInnerLeftRow, error) {
	rows, err := q.db.QueryContext(ctx, testInnerLeftInnerLeft)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TestInnerLeftInnerLeftRow
	for rows.Next() {
		var i TestInnerLeftInnerLeftRow
		if err := rows.Scan(
			&i.A,
			&i.B,
			&i.C,
			&i.D,
			&i.E,
		); err != nil {
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

const testLeftInner = `-- name: TestLeftInner :many
SELECT a.a, b.b, c.c
FROM a
LEFT JOIN b ON b.a_id = a.id
INNER JOIN c ON c.a_id = a.id
`

type TestLeftInnerRow struct {
	A string
	B sql.NullString
	C string
}

func (q *Queries) TestLeftInner(ctx context.Context) ([]TestLeftInnerRow, error) {
	rows, err := q.db.QueryContext(ctx, testLeftInner)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TestLeftInnerRow
	for rows.Next() {
		var i TestLeftInnerRow
		if err := rows.Scan(&i.A, &i.B, &i.C); err != nil {
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

const testLeftInnerLeftInner = `-- name: TestLeftInnerLeftInner :many
SELECT a.a, b.b, c.c, d.d, e.e
FROM a
LEFT JOIN b ON b.a_id = a.id
INNER JOIN c ON c.a_id = a.id
LEFT JOIN d ON d.a_id = a.id
INNER JOIN e ON e.a_id = a.id
`

type TestLeftInnerLeftInnerRow struct {
	A string
	B sql.NullString
	C string
	D sql.NullString
	E string
}

func (q *Queries) TestLeftInnerLeftInner(ctx context.Context) ([]TestLeftInnerLeftInnerRow, error) {
	rows, err := q.db.QueryContext(ctx, testLeftInnerLeftInner)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TestLeftInnerLeftInnerRow
	for rows.Next() {
		var i TestLeftInnerLeftInnerRow
		if err := rows.Scan(
			&i.A,
			&i.B,
			&i.C,
			&i.D,
			&i.E,
		); err != nil {
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
