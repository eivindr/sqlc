// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package test

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type DebugCenum string

const (
	DebugCenumOne   DebugCenum = "one"
	DebugCenumTwo   DebugCenum = "two"
	DebugCenumThree DebugCenum = "three"
)

func (e *DebugCenum) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = DebugCenum(s)
	case string:
		*e = DebugCenum(s)
	default:
		return fmt.Errorf("unsupported scan type for DebugCenum: %T", src)
	}
	return nil
}

type NullDebugCenum struct {
	DebugCenum DebugCenum
	Valid      bool // Valid is true if DebugCenum is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullDebugCenum) Scan(value interface{}) error {
	if value == nil {
		ns.DebugCenum, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.DebugCenum.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullDebugCenum) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.DebugCenum), nil
}

type DebugCset string

const (
	DebugCsetOne   DebugCset = "one"
	DebugCsetTwo   DebugCset = "two"
	DebugCsetThree DebugCset = "three"
)

func (e *DebugCset) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = DebugCset(s)
	case string:
		*e = DebugCset(s)
	default:
		return fmt.Errorf("unsupported scan type for DebugCset: %T", src)
	}
	return nil
}

type NullDebugCset struct {
	DebugCset DebugCset
	Valid     bool // Valid is true if DebugCset is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullDebugCset) Scan(value interface{}) error {
	if value == nil {
		ns.DebugCset, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.DebugCset.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullDebugCset) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.DebugCset), nil
}

type Debug struct {
	ID               int64
	Csmallint        int32
	Cint             int32
	Cinteger         int32
	Cdecimal         string
	Cnumeric         string
	Cfloat           float64
	Creal            float64
	Cdoubleprecision float64
	Cdouble          float64
	Cdec             string
	Cfixed           string
	Ctinyint         int32
	Cbool            bool
	Cmediumint       int32
	Cbit             interface{}
	Cdate            time.Time
	Cdatetime        time.Time
	Ctimestamp       time.Time
	Ctime            time.Time
	Cyear            int32
	Cchar            string
	Cvarchar         string
	Cbinary          []byte
	Cvarbinary       []byte
	Ctinyblob        []byte
	Cblob            []byte
	Cmediumblob      []byte
	Clongblob        []byte
	Ctinytext        string
	Ctext            string
	Cmediumtext      string
	Clongtext        string
	Cenum            NullDebugCenum
	Cset             DebugCset
	Cjson            json.RawMessage
}
