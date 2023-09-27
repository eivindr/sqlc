// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package querytest

import (
	"database/sql/driver"
	"fmt"
)

type Digit string

const (
	Digit0       Digit = "0"
	Digit1       Digit = "1"
	Digit2       Digit = "2"
	Digit3       Digit = "3"
	Digit4       Digit = "4"
	Digit5       Digit = "5"
	Digit6       Digit = "6"
	Digit7       Digit = "7"
	Digit8       Digit = "8"
	Digit9       Digit = "9"
	DigitValue10 Digit = "#"
	DigitValue11 Digit = "*"
)

func (e *Digit) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Digit(s)
	case string:
		*e = Digit(s)
	default:
		return fmt.Errorf("unsupported scan type for Digit: %T", src)
	}
	return nil
}

type NullDigit struct {
	Digit Digit
	Valid bool // Valid is true if Digit is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullDigit) Scan(value interface{}) error {
	if value == nil {
		ns.Digit, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Digit.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullDigit) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Digit), nil
}

type Foobar string

const (
	FoobarFooA Foobar = "foo-a"
	FoobarFooB Foobar = "foo_b"
	FoobarFooC Foobar = "foo:c"
	FoobarFooD Foobar = "foo/d"
	FoobarFooe Foobar = "foo@e"
	FoobarFoof Foobar = "foo+f"
	FoobarFoog Foobar = "foo!g"
)

func (e *Foobar) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Foobar(s)
	case string:
		*e = Foobar(s)
	default:
		return fmt.Errorf("unsupported scan type for Foobar: %T", src)
	}
	return nil
}

type NullFoobar struct {
	Foobar Foobar
	Valid  bool // Valid is true if Foobar is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullFoobar) Scan(value interface{}) error {
	if value == nil {
		ns.Foobar, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Foobar.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullFoobar) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Foobar), nil
}

type Foo struct {
	Val Foobar
}
