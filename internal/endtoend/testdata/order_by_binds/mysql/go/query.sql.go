// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: query.sql

package querytest

import (
	"context"
)

const listAuthorsColumnSort = `-- name: ListAuthorsColumnSort :many
SELECT  id, name, bio FROM authors
WHERE   id > ? 
ORDER   BY CASE WHEN ? = 'name' THEN name END
`

type ListAuthorsColumnSortParams struct {
	MinID      int64
	SortColumn interface{}
}

func (q *Queries) ListAuthorsColumnSort(ctx context.Context, arg ListAuthorsColumnSortParams) ([]Author, error) {
	rows, err := q.db.QueryContext(ctx, listAuthorsColumnSort, arg.MinID, arg.SortColumn)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Author
	for rows.Next() {
		var i Author
		if err := rows.Scan(&i.ID, &i.Name, &i.Bio); err != nil {
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

const listAuthorsNameSort = `-- name: ListAuthorsNameSort :many
SELECT  id, name, bio FROM authors
WHERE   id > ?
ORDER   BY name ASC
`

func (q *Queries) ListAuthorsNameSort(ctx context.Context, minID int64) ([]Author, error) {
	rows, err := q.db.QueryContext(ctx, listAuthorsNameSort, minID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Author
	for rows.Next() {
		var i Author
		if err := rows.Scan(&i.ID, &i.Name, &i.Bio); err != nil {
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
