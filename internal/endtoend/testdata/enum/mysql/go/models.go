// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package querytest

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
)

type UsersShirtSize string

const (
	UsersShirtSizeXSmall UsersShirtSize = "x-small"
	UsersShirtSizeSmall  UsersShirtSize = "small"
	UsersShirtSizeMedium UsersShirtSize = "medium"
	UsersShirtSizeLarge  UsersShirtSize = "large"
	UsersShirtSizeXLarge UsersShirtSize = "x-large"
)

func (e *UsersShirtSize) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UsersShirtSize(s)
	case string:
		*e = UsersShirtSize(s)
	default:
		return fmt.Errorf("unsupported scan type for UsersShirtSize: %T", src)
	}
	return nil
}

type NullUsersShirtSize struct {
	UsersShirtSize UsersShirtSize
	Valid          bool // Valid is true if UsersShirtSize is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUsersShirtSize) Scan(value interface{}) error {
	if value == nil {
		ns.UsersShirtSize, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UsersShirtSize.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUsersShirtSize) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.UsersShirtSize), nil
}

type UsersShoeSize string

const (
	UsersShoeSizeXSmall UsersShoeSize = "x-small"
	UsersShoeSizeSmall  UsersShoeSize = "small"
	UsersShoeSizeMedium UsersShoeSize = "medium"
	UsersShoeSizeLarge  UsersShoeSize = "large"
	UsersShoeSizeXLarge UsersShoeSize = "x-large"
)

func (e *UsersShoeSize) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UsersShoeSize(s)
	case string:
		*e = UsersShoeSize(s)
	default:
		return fmt.Errorf("unsupported scan type for UsersShoeSize: %T", src)
	}
	return nil
}

type NullUsersShoeSize struct {
	UsersShoeSize UsersShoeSize
	Valid         bool // Valid is true if UsersShoeSize is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUsersShoeSize) Scan(value interface{}) error {
	if value == nil {
		ns.UsersShoeSize, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UsersShoeSize.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUsersShoeSize) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.UsersShoeSize), nil
}

type User struct {
	ID        int32
	FirstName string
	LastName  sql.NullString
	Age       int32
	ShoeSize  UsersShoeSize
	ShirtSize NullUsersShirtSize
}
