// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package querytest

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type CaseIntent struct {
	ID               int32
	CaseIntentString string
	Description      string
	Author           string
}

type CaseIntentParentJoin struct {
	CaseIntentID       int64
	CaseIntentParentID int64
}

type CaseIntentVersion struct {
	VersionID int32
	Reviewer  string
	CreatedAt pgtype.Timestamptz
}

type CaseIntentVersionJoin struct {
	CaseIntentID        int64
	CaseIntentVersionID int32
}
