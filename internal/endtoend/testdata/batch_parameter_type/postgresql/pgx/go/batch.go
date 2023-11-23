// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: batch.go

package querytest

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

var (
	ErrBatchAlreadyClosed = errors.New("batch already closed")
)

const insertMappping = `-- name: InsertMappping :batchexec
WITH
    table1
        AS (
            SELECT
                version
            FROM
                solar_commcard_mapping
            WHERE
                "deviceId" = $1
            ORDER BY
                "updatedAt" DESC
            LIMIT
                1
        )
INSERT
INTO
    solar_commcard_mapping
        ("deviceId", version, sn, "updatedAt")
SELECT
    $1, $2::text, $3, $4
WHERE
    NOT
        EXISTS(
            SELECT
                version
            FROM
                table1
            WHERE
                table1.version = $2::text
        )
    OR NOT EXISTS(SELECT version FROM table1)
`

type InsertMapppingBatchResults struct {
	br     pgx.BatchResults
	tot    int
	closed bool
}

type InsertMapppingParams struct {
	DeviceId  int64
	Version   string
	Sn        string
	UpdatedAt pgtype.Timestamptz
}

func (q *Queries) InsertMappping(ctx context.Context, arg []InsertMapppingParams) *InsertMapppingBatchResults {
	batch := &pgx.Batch{}
	for _, a := range arg {
		vals := []interface{}{
			a.DeviceId,
			a.Version,
			a.Sn,
			a.UpdatedAt,
		}
		batch.Queue(insertMappping, vals...)
	}
	br := q.db.SendBatch(ctx, batch)
	return &InsertMapppingBatchResults{br, len(arg), false}
}

func (b *InsertMapppingBatchResults) Exec(f func(int, error)) {
	defer b.br.Close()
	for t := 0; t < b.tot; t++ {
		if b.closed {
			if f != nil {
				f(t, ErrBatchAlreadyClosed)
			}
			continue
		}
		_, err := b.br.Exec()
		if f != nil {
			f(t, err)
		}
	}
}

func (b *InsertMapppingBatchResults) Close() error {
	b.closed = true
	return b.br.Close()
}
