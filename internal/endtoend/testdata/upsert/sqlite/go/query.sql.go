// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: query.sql

package querytest

import (
	"context"
)

const upsertLocation = `-- name: UpsertLocation :exec
INSERT INTO locations (
    name,
    address,
    zip_code,
    latitude,
    longitude
)
VALUES (?, ?, ?, ?, ?)
ON CONFLICT(name) DO UPDATE SET 
    name = excluded.name,
    address = excluded.address,
    zip_code = excluded.zip_code,
    latitude = excluded.latitude,
    longitude = excluded.longitude
`

type UpsertLocationParams struct {
	Name      string
	Address   string
	ZipCode   int64
	Latitude  float64
	Longitude float64
}

func (q *Queries) UpsertLocation(ctx context.Context, arg UpsertLocationParams) error {
	_, err := q.db.ExecContext(ctx, upsertLocation,
		arg.Name,
		arg.Address,
		arg.ZipCode,
		arg.Latitude,
		arg.Longitude,
	)
	return err
}
