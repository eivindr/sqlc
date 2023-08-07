// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package querytest

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	CreateMemories(ctx context.Context, vampireID []pgtype.UUID) ([]Memory, error)
	GetVampireIDs(ctx context.Context, vampireID []pgtype.UUID) ([]pgtype.UUID, error)
}

var _ Querier = (*Queries)(nil)
