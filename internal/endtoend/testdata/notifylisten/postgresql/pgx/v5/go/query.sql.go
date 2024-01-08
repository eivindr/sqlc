// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package querytest

import (
	"context"
)

const listenTest = `-- name: ListenTest :exec
LISTEN test
`

func (q *Queries) ListenTest(ctx context.Context) error {
	_, err := q.db.Exec(ctx, listenTest)
	return err
}

const notifyTest = `-- name: NotifyTest :exec
NOTIFY test
`

func (q *Queries) NotifyTest(ctx context.Context) error {
	_, err := q.db.Exec(ctx, notifyTest)
	return err
}

const notifyWithMessage = `-- name: NotifyWithMessage :exec
NOTIFY test, 'msg'
`

func (q *Queries) NotifyWithMessage(ctx context.Context) error {
	_, err := q.db.Exec(ctx, notifyWithMessage)
	return err
}
