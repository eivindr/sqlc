// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: query.sql

package querytest

import (
	"context"
)

const addAuthor = `-- name: AddAuthor :execlastid
INSERT INTO authors (
    address,
    name,
    bio
) VALUES (
    ?,
    COALESCE(?, ""),
    COALESCE(?, "")
)
`

type AddAuthorParams struct {
	Address        string
	CalName        interface{}
	CalDescription interface{}
}

func (q *Queries) AddAuthor(ctx context.Context, arg AddAuthorParams) (int64, error) {
	result, err := q.db.ExecContext(ctx, addAuthor, arg.Address, arg.CalName, arg.CalDescription)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

const addEvent = `-- name: AddEvent :execlastid
INSERT INTO ` + "`" + `Event` + "`" + ` (
    Timezone
) VALUES (
    (CASE WHEN ? = "calendar" THEN (SELECT cal.Timezone FROM Calendar cal WHERE cal.IdKey = ?) ELSE ? END)
)
`

type AddEventParams struct {
	Timezone      interface{}
	CalendarIdKey string
}

func (q *Queries) AddEvent(ctx context.Context, arg AddEventParams) (int64, error) {
	result, err := q.db.ExecContext(ctx, addEvent, arg.Timezone, arg.CalendarIdKey, arg.Timezone)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}
