// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package querytest

import (
	"database/sql"

	"github.com/google/uuid"
)

type Foo struct {
	Description sql.NullString
	Bar         uuid.NullUUID
	Baz         uuid.UUID
}
