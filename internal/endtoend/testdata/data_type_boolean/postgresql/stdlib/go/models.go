// Code generated by sqlc. DO NOT EDIT.

package querytest

import (
	"database/sql"
)

type Bar struct {
	ColA sql.NullBool
	ColB sql.NullBool
}

type Foo struct {
	ColA bool
	ColB bool
}
