// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package querytest

import (
	"context"
	"database/sql"
)

const subquery = `-- name: Subquery :many
SELECT 
	a,
	name,
	(SELECT alias FROM bar WHERE bar.a=foo.a AND alias = $1 ORDER BY bar.a DESC limit 1) as alias
FROM FOO WHERE a = $2
`

type SubqueryParams struct {
	Column1 sql.NullString
	Column2 sql.NullInt32
}

type SubqueryRow struct {
	A     int32
	Name  sql.NullString
	Alias sql.NullString
}

func (q *Queries) Subquery(ctx context.Context, arg SubqueryParams) ([]SubqueryRow, error) {
	rows, err := q.db.QueryContext(ctx, subquery, arg.Column1, arg.Column2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SubqueryRow
	for rows.Next() {
		var i SubqueryRow
		if err := rows.Scan(&i.A, &i.Name, &i.Alias); err != nil {
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
