// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: query.sql

package querytest

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const fooByAandB = `-- name: FooByAandB :many
SELECT a, b FROM foo 
WHERE a = $1 and b = $2
`

type FooByAandBParams struct {
	A pgtype.Text
	B pgtype.Text
}

func (q *Queries) FooByAandB(ctx context.Context, arg FooByAandBParams) ([]Foo, error) {
	rows, err := q.db.Query(ctx, fooByAandB, arg.A, arg.B)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Foo
	for rows.Next() {
		var i Foo
		if err := rows.Scan(&i.A, &i.B); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
