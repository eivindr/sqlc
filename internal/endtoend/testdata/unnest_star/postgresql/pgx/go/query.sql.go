// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package querytest

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getPlanItems = `-- name: GetPlanItems :many
SELECT p.plan_id, p.item_id
FROM (SELECT unnest FROM unnest($1::bigint[])) AS i(req_item_id),
LATERAL (
    SELECT plan_id, item_id
    FROM plan_items
    WHERE
        item_id = i.req_item_id AND
        ($2 = 0 OR plan_id < $2)
    ORDER BY plan_id DESC
    LIMIT $3
) p
`

type GetPlanItemsParams struct {
	Ids        []int64
	After      pgtype.Int4
	LimitCount int64
}

type GetPlanItemsRow struct {
	PlanID int64
	ItemID int64
}

func (q *Queries) GetPlanItems(ctx context.Context, arg GetPlanItemsParams) ([]GetPlanItemsRow, error) {
	rows, err := q.db.Query(ctx, getPlanItems, arg.Ids, arg.After, arg.LimitCount)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPlanItemsRow
	for rows.Next() {
		var i GetPlanItemsRow
		if err := rows.Scan(&i.PlanID, &i.ItemID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
