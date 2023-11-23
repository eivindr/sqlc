// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: query.sql

package querytest

import (
	"context"
	"strings"
)

const funcParamIdent = `-- name: FuncParamIdent :many
SELECT name FROM foo
WHERE name = ?1
  AND id IN (/*SLICE:favourites*/?)
`

type FuncParamIdentParams struct {
	Slug       string
	Favourites []int64
}

func (q *Queries) FuncParamIdent(ctx context.Context, arg FuncParamIdentParams) ([]string, error) {
	query := funcParamIdent
	var queryParams []interface{}
	queryParams = append(queryParams, arg.Slug)
	if len(arg.Favourites) > 0 {
		for _, v := range arg.Favourites {
			queryParams = append(queryParams, v)
		}
		query = strings.Replace(query, "/*SLICE:favourites*/?", strings.Repeat(",?", len(arg.Favourites))[1:], 1)
	} else {
		query = strings.Replace(query, "/*SLICE:favourites*/?", "NULL", 1)
	}
	rows, err := q.query(ctx, nil, query, queryParams...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		items = append(items, name)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
