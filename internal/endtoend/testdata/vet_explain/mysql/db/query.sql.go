// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package test

import (
	"context"
	"encoding/json"
	"time"
)

const selectByCbinary = `-- name: SelectByCbinary :one
SELECT id FROM debug
WHERE Cbinary = ? LIMIT 1
`

func (q *Queries) SelectByCbinary(ctx context.Context, cbinary []byte) (int64, error) {
	row := q.db.QueryRowContext(ctx, selectByCbinary, cbinary)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const selectByCbit = `-- name: SelectByCbit :one
SELECT id FROM debug
WHERE Cbit = ? LIMIT 1
`

func (q *Queries) SelectByCbit(ctx context.Context, cbit interface{}) (int64, error) {
	row := q.db.QueryRowContext(ctx, selectByCbit, cbit)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const selectByCblob = `-- name: SelectByCblob :one
SELECT id FROM debug
WHERE Cblob = ? LIMIT 1
`

func (q *Queries) SelectByCblob(ctx context.Context, cblob []byte) (int64, error) {
	row := q.db.QueryRowContext(ctx, selectByCblob, cblob)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const selectByCbool = `-- name: SelectByCbool :one
SELECT id FROM debug
WHERE Cbool = ? LIMIT 1
`

func (q *Queries) SelectByCbool(ctx context.Context, cbool bool) (int64, error) {
	row := q.db.QueryRowContext(ctx, selectByCbool, cbool)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const selectByCchar = `-- name: SelectByCchar :one
SELECT id FROM debug
WHERE Cchar = ? LIMIT 1
`

func (q *Queries) SelectByCchar(ctx context.Context, cchar string) (int64, error) {
	row := q.db.QueryRowContext(ctx, selectByCchar, cchar)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const selectByCdate = `-- name: SelectByCdate :one
SELECT id FROM debug
WHERE Cdate = ? LIMIT 1
`

func (q *Queries) SelectByCdate(ctx context.Context, cdate time.Time) (int64, error) {
	row := q.db.QueryRowContext(ctx, selectByCdate, cdate)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const selectByCdatetime = `-- name: SelectByCdatetime :one
SELECT id FROM debug
WHERE Cdatetime = ? LIMIT 1
`

func (q *Queries) SelectByCdatetime(ctx context.Context, cdatetime time.Time) (int64, error) {
	row := q.db.QueryRowContext(ctx, selectByCdatetime, cdatetime)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const selectByCdec = `-- name: SelectByCdec :one
SELECT id FROM debug
WHERE Cdec = ? LIMIT 1
`

func (q *Queries) SelectByCdec(ctx context.Context, cdec string) (int64, error) {
	row := q.db.QueryRowContext(ctx, selectByCdec, cdec)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const selectByCdecimal = `-- name: SelectByCdecimal :one
SELECT id FROM debug
WHERE Cdecimal = ? LIMIT 1
`

func (q *Queries) SelectByCdecimal(ctx context.Context, cdecimal string) (int64, error) {
	row := q.db.QueryRowContext(ctx, selectByCdecimal, cdecimal)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const selectByCdouble = `-- name: SelectByCdouble :one
SELECT id FROM debug
WHERE Cdouble = ? LIMIT 1
`

func (q *Queries) SelectByCdouble(ctx context.Context, cdouble float64) (int64, error) {
	row := q.db.QueryRowContext(ctx, selectByCdouble, cdouble)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const selectByCdoubleprecision = `-- name: SelectByCdoubleprecision :one
SELECT id FROM debug
WHERE Cdoubleprecision = ? LIMIT 1
`

func (q *Queries) SelectByCdoubleprecision(ctx context.Context, cdoubleprecision float64) (int64, error) {
	row := q.db.QueryRowContext(ctx, selectByCdoubleprecision, cdoubleprecision)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const selectByCenum = `-- name: SelectByCenum :one
SELECT id FROM debug
WHERE Cenum = ? LIMIT 1
`

func (q *Queries) SelectByCenum(ctx context.Context, cenum NullDebugCenum) (int64, error) {
	row := q.db.QueryRowContext(ctx, selectByCenum, cenum)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const selectByCfixed = `-- name: SelectByCfixed :one
SELECT id FROM debug
WHERE Cfixed = ? LIMIT 1
`

func (q *Queries) SelectByCfixed(ctx context.Context, cfixed string) (int64, error) {
	row := q.db.QueryRowContext(ctx, selectByCfixed, cfixed)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const selectByCfloat = `-- name: SelectByCfloat :one
SELECT id FROM debug
WHERE Cfloat = ? LIMIT 1
`

func (q *Queries) SelectByCfloat(ctx context.Context, cfloat float64) (int64, error) {
	row := q.db.QueryRowContext(ctx, selectByCfloat, cfloat)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const selectByCint = `-- name: SelectByCint :one
SELECT id FROM debug
WHERE Cint = ? LIMIT 1
`

func (q *Queries) SelectByCint(ctx context.Context, cint int32) (int64, error) {
	row := q.db.QueryRowContext(ctx, selectByCint, cint)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const selectByCinteger = `-- name: SelectByCinteger :one
SELECT id FROM debug
WHERE Cinteger = ? LIMIT 1
`

func (q *Queries) SelectByCinteger(ctx context.Context, cinteger int32) (int64, error) {
	row := q.db.QueryRowContext(ctx, selectByCinteger, cinteger)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const selectByCjson = `-- name: SelectByCjson :one
SELECT id FROM debug
WHERE Cjson = ? LIMIT 1
`

func (q *Queries) SelectByCjson(ctx context.Context, cjson json.RawMessage) (int64, error) {
	row := q.db.QueryRowContext(ctx, selectByCjson, cjson)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const selectByClongblob = `-- name: SelectByClongblob :one
SELECT id FROM debug
WHERE Clongblob = ? LIMIT 1
`

func (q *Queries) SelectByClongblob(ctx context.Context, clongblob []byte) (int64, error) {
	row := q.db.QueryRowContext(ctx, selectByClongblob, clongblob)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const selectByClongtext = `-- name: SelectByClongtext :one
SELECT id FROM debug
WHERE Clongtext = ? LIMIT 1
`

func (q *Queries) SelectByClongtext(ctx context.Context, clongtext string) (int64, error) {
	row := q.db.QueryRowContext(ctx, selectByClongtext, clongtext)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const selectByCmediumblob = `-- name: SelectByCmediumblob :one
SELECT id FROM debug
WHERE Cmediumblob = ? LIMIT 1
`

func (q *Queries) SelectByCmediumblob(ctx context.Context, cmediumblob []byte) (int64, error) {
	row := q.db.QueryRowContext(ctx, selectByCmediumblob, cmediumblob)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const selectByCmediumint = `-- name: SelectByCmediumint :one
SELECT id FROM debug
WHERE Cmediumint = ? LIMIT 1
`

func (q *Queries) SelectByCmediumint(ctx context.Context, cmediumint int32) (int64, error) {
	row := q.db.QueryRowContext(ctx, selectByCmediumint, cmediumint)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const selectByCmediumtext = `-- name: SelectByCmediumtext :one
SELECT id FROM debug
WHERE Cmediumtext = ? LIMIT 1
`

func (q *Queries) SelectByCmediumtext(ctx context.Context, cmediumtext string) (int64, error) {
	row := q.db.QueryRowContext(ctx, selectByCmediumtext, cmediumtext)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const selectByCnumeric = `-- name: SelectByCnumeric :one
SELECT id FROM debug
WHERE Cnumeric = ? LIMIT 1
`

func (q *Queries) SelectByCnumeric(ctx context.Context, cnumeric string) (int64, error) {
	row := q.db.QueryRowContext(ctx, selectByCnumeric, cnumeric)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const selectByCreal = `-- name: SelectByCreal :one
SELECT id FROM debug
WHERE Creal = ? LIMIT 1
`

func (q *Queries) SelectByCreal(ctx context.Context, creal float64) (int64, error) {
	row := q.db.QueryRowContext(ctx, selectByCreal, creal)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const selectByCset = `-- name: SelectByCset :one
SELECT id FROM debug
WHERE Cset = ? LIMIT 1
`

func (q *Queries) SelectByCset(ctx context.Context, cset DebugCset) (int64, error) {
	row := q.db.QueryRowContext(ctx, selectByCset, cset)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const selectByCsmallint = `-- name: SelectByCsmallint :one
SELECT id FROM debug
WHERE Csmallint = ? LIMIT 1
`

func (q *Queries) SelectByCsmallint(ctx context.Context, csmallint int16) (int64, error) {
	row := q.db.QueryRowContext(ctx, selectByCsmallint, csmallint)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const selectByCtext = `-- name: SelectByCtext :one
SELECT id FROM debug
WHERE Ctext = ? LIMIT 1
`

func (q *Queries) SelectByCtext(ctx context.Context, ctext string) (int64, error) {
	row := q.db.QueryRowContext(ctx, selectByCtext, ctext)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const selectByCtime = `-- name: SelectByCtime :one
SELECT id FROM debug
WHERE Ctime = ? LIMIT 1
`

func (q *Queries) SelectByCtime(ctx context.Context, ctime time.Time) (int64, error) {
	row := q.db.QueryRowContext(ctx, selectByCtime, ctime)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const selectByCtimestamp = `-- name: SelectByCtimestamp :one
SELECT id FROM debug
WHERE Ctimestamp = ? LIMIT 1
`

func (q *Queries) SelectByCtimestamp(ctx context.Context, ctimestamp time.Time) (int64, error) {
	row := q.db.QueryRowContext(ctx, selectByCtimestamp, ctimestamp)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const selectByCtinyblob = `-- name: SelectByCtinyblob :one
SELECT id FROM debug
WHERE Ctinyblob = ? LIMIT 1
`

func (q *Queries) SelectByCtinyblob(ctx context.Context, ctinyblob []byte) (int64, error) {
	row := q.db.QueryRowContext(ctx, selectByCtinyblob, ctinyblob)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const selectByCtinyint = `-- name: SelectByCtinyint :one
SELECT id FROM debug
WHERE Ctinyint = ? LIMIT 1
`

func (q *Queries) SelectByCtinyint(ctx context.Context, ctinyint int8) (int64, error) {
	row := q.db.QueryRowContext(ctx, selectByCtinyint, ctinyint)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const selectByCtinytext = `-- name: SelectByCtinytext :one
SELECT id FROM debug
WHERE Ctinytext = ? LIMIT 1
`

func (q *Queries) SelectByCtinytext(ctx context.Context, ctinytext string) (int64, error) {
	row := q.db.QueryRowContext(ctx, selectByCtinytext, ctinytext)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const selectByCvarbinary = `-- name: SelectByCvarbinary :one
SELECT id FROM debug
WHERE Cvarbinary = ? LIMIT 1
`

func (q *Queries) SelectByCvarbinary(ctx context.Context, cvarbinary []byte) (int64, error) {
	row := q.db.QueryRowContext(ctx, selectByCvarbinary, cvarbinary)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const selectByCvarchar = `-- name: SelectByCvarchar :one
SELECT id FROM debug
WHERE Cvarchar = ? LIMIT 1
`

func (q *Queries) SelectByCvarchar(ctx context.Context, cvarchar string) (int64, error) {
	row := q.db.QueryRowContext(ctx, selectByCvarchar, cvarchar)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const selectByCyear = `-- name: SelectByCyear :one
SELECT id FROM debug
WHERE Cyear = ? LIMIT 1
`

func (q *Queries) SelectByCyear(ctx context.Context, cyear int16) (int64, error) {
	row := q.db.QueryRowContext(ctx, selectByCyear, cyear)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const selectById = `-- name: SelectById :one
SELECT id FROM debug
WHERE id = ? LIMIT 1
`

func (q *Queries) SelectById(ctx context.Context, id int64) (int64, error) {
	row := q.db.QueryRowContext(ctx, selectById, id)
	err := row.Scan(&id)
	return id, err
}
