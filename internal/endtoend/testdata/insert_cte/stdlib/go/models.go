// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package querytest

import (
	"database/sql"
	"time"
)

type Td3Code struct {
	ID        int32
	TsCreated time.Time
	TsUpdated time.Time
	CreatedBy string
	UpdatedBy string
	Code      sql.NullString
	Hash      sql.NullString
	IsPrivate sql.NullBool
}

type Td3TestCode struct {
	ID        int32
	TsCreated time.Time
	TsUpdated time.Time
	CreatedBy string
	UpdatedBy string
	TestID    int32
	CodeHash  string
}
