// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package querytest

import (
	"database/sql"

	"github.com/google/uuid"
)

type Llc struct {
	PartyID           uuid.UUID
	Name              string
	LegalName         string
	IncorporationDate sql.NullTime
}

type Organisation struct {
	PartyID   uuid.UUID
	Name      string
	LegalName string
}

type Party struct {
	PartyID uuid.UUID
	Name    string
}

type Person struct {
	PartyID   uuid.UUID
	Name      string
	FirstName string
	LastName  string
}
