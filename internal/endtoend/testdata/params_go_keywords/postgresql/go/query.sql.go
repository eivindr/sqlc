// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: query.sql

package querytest

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const keywordBreak = `-- name: KeywordBreak :exec
SELECT $1::text
`

func (q *Queries) KeywordBreak(ctx context.Context, break_ string) error {
	_, err := q.db.Exec(ctx, keywordBreak, break_)
	return err
}

const keywordCase = `-- name: KeywordCase :exec
SELECT $1::text
`

func (q *Queries) KeywordCase(ctx context.Context, case_ string) error {
	_, err := q.db.Exec(ctx, keywordCase, case_)
	return err
}

const keywordChan = `-- name: KeywordChan :exec
SELECT $1::text
`

func (q *Queries) KeywordChan(ctx context.Context, chan_ string) error {
	_, err := q.db.Exec(ctx, keywordChan, chan_)
	return err
}

const keywordConst = `-- name: KeywordConst :exec
SELECT $1::text
`

func (q *Queries) KeywordConst(ctx context.Context, const_ string) error {
	_, err := q.db.Exec(ctx, keywordConst, const_)
	return err
}

const keywordContinue = `-- name: KeywordContinue :exec
SELECT $1::text
`

func (q *Queries) KeywordContinue(ctx context.Context, continue_ string) error {
	_, err := q.db.Exec(ctx, keywordContinue, continue_)
	return err
}

const keywordDefault = `-- name: KeywordDefault :exec
SELECT $1::text
`

func (q *Queries) KeywordDefault(ctx context.Context, default_ string) error {
	_, err := q.db.Exec(ctx, keywordDefault, default_)
	return err
}

const keywordDefer = `-- name: KeywordDefer :exec
SELECT $1::text
`

func (q *Queries) KeywordDefer(ctx context.Context, defer_ string) error {
	_, err := q.db.Exec(ctx, keywordDefer, defer_)
	return err
}

const keywordElse = `-- name: KeywordElse :exec
SELECT $1::text
`

func (q *Queries) KeywordElse(ctx context.Context, else_ string) error {
	_, err := q.db.Exec(ctx, keywordElse, else_)
	return err
}

const keywordFallthrough = `-- name: KeywordFallthrough :exec
SELECT $1::text
`

func (q *Queries) KeywordFallthrough(ctx context.Context, fallthrough_ string) error {
	_, err := q.db.Exec(ctx, keywordFallthrough, fallthrough_)
	return err
}

const keywordFor = `-- name: KeywordFor :exec
SELECT $1::text
`

func (q *Queries) KeywordFor(ctx context.Context, for_ string) error {
	_, err := q.db.Exec(ctx, keywordFor, for_)
	return err
}

const keywordFunc = `-- name: KeywordFunc :exec
SELECT $1::text
`

func (q *Queries) KeywordFunc(ctx context.Context, func_ string) error {
	_, err := q.db.Exec(ctx, keywordFunc, func_)
	return err
}

const keywordGo = `-- name: KeywordGo :exec
SELECT $1::text
`

func (q *Queries) KeywordGo(ctx context.Context, go_ string) error {
	_, err := q.db.Exec(ctx, keywordGo, go_)
	return err
}

const keywordGoto = `-- name: KeywordGoto :exec
SELECT $1::text
`

func (q *Queries) KeywordGoto(ctx context.Context, goto_ string) error {
	_, err := q.db.Exec(ctx, keywordGoto, goto_)
	return err
}

const keywordIf = `-- name: KeywordIf :exec
SELECT $1::text
`

func (q *Queries) KeywordIf(ctx context.Context, if_ string) error {
	_, err := q.db.Exec(ctx, keywordIf, if_)
	return err
}

const keywordImport = `-- name: KeywordImport :exec
SELECT $1::text
`

func (q *Queries) KeywordImport(ctx context.Context, import_ string) error {
	_, err := q.db.Exec(ctx, keywordImport, import_)
	return err
}

const keywordInterface = `-- name: KeywordInterface :exec
SELECT $1::text
`

func (q *Queries) KeywordInterface(ctx context.Context, interface_ string) error {
	_, err := q.db.Exec(ctx, keywordInterface, interface_)
	return err
}

const keywordMap = `-- name: KeywordMap :exec
SELECT $1::text
`

func (q *Queries) KeywordMap(ctx context.Context, map_ string) error {
	_, err := q.db.Exec(ctx, keywordMap, map_)
	return err
}

const keywordPackage = `-- name: KeywordPackage :exec
SELECT $1::text
`

func (q *Queries) KeywordPackage(ctx context.Context, package_ string) error {
	_, err := q.db.Exec(ctx, keywordPackage, package_)
	return err
}

const keywordRange = `-- name: KeywordRange :exec
SELECT $1::text
`

func (q *Queries) KeywordRange(ctx context.Context, range_ string) error {
	_, err := q.db.Exec(ctx, keywordRange, range_)
	return err
}

const keywordReturn = `-- name: KeywordReturn :exec
SELECT $1::text
`

func (q *Queries) KeywordReturn(ctx context.Context, return_ string) error {
	_, err := q.db.Exec(ctx, keywordReturn, return_)
	return err
}

const keywordSelect = `-- name: KeywordSelect :exec
SELECT $1::text
`

func (q *Queries) KeywordSelect(ctx context.Context, select_ string) error {
	_, err := q.db.Exec(ctx, keywordSelect, select_)
	return err
}

const keywordStruct = `-- name: KeywordStruct :exec
SELECT $1::text
`

func (q *Queries) KeywordStruct(ctx context.Context, struct_ string) error {
	_, err := q.db.Exec(ctx, keywordStruct, struct_)
	return err
}

const keywordSwitch = `-- name: KeywordSwitch :exec
SELECT $1::text
`

func (q *Queries) KeywordSwitch(ctx context.Context, switch_ string) error {
	_, err := q.db.Exec(ctx, keywordSwitch, switch_)
	return err
}

const keywordType = `-- name: KeywordType :exec
SELECT $1::text
`

func (q *Queries) KeywordType(ctx context.Context, type_ string) error {
	_, err := q.db.Exec(ctx, keywordType, type_)
	return err
}

const keywordVar = `-- name: KeywordVar :exec
SELECT $1::text
`

func (q *Queries) KeywordVar(ctx context.Context, var_ string) error {
	_, err := q.db.Exec(ctx, keywordVar, var_)
	return err
}

const selectBreak = `-- name: SelectBreak :one
SELECT "break" FROM go_keywords
`

func (q *Queries) SelectBreak(ctx context.Context) (pgtype.Text, error) {
	row := q.db.QueryRow(ctx, selectBreak)
	var break_ pgtype.Text
	err := row.Scan(&break_)
	return break_, err
}

const selectCase = `-- name: SelectCase :one
SELECT "case" FROM go_keywords
`

func (q *Queries) SelectCase(ctx context.Context) (pgtype.Text, error) {
	row := q.db.QueryRow(ctx, selectCase)
	var case_ pgtype.Text
	err := row.Scan(&case_)
	return case_, err
}

const selectChan = `-- name: SelectChan :one
SELECT "chan" FROM go_keywords
`

func (q *Queries) SelectChan(ctx context.Context) (pgtype.Text, error) {
	row := q.db.QueryRow(ctx, selectChan)
	var chan_ pgtype.Text
	err := row.Scan(&chan_)
	return chan_, err
}

const selectConst = `-- name: SelectConst :one
SELECT "const" FROM go_keywords
`

func (q *Queries) SelectConst(ctx context.Context) (pgtype.Text, error) {
	row := q.db.QueryRow(ctx, selectConst)
	var const_ pgtype.Text
	err := row.Scan(&const_)
	return const_, err
}

const selectContinue = `-- name: SelectContinue :one
SELECT "continue" FROM go_keywords
`

func (q *Queries) SelectContinue(ctx context.Context) (pgtype.Text, error) {
	row := q.db.QueryRow(ctx, selectContinue)
	var continue_ pgtype.Text
	err := row.Scan(&continue_)
	return continue_, err
}

const selectDefault = `-- name: SelectDefault :one
SELECT "default" FROM go_keywords
`

func (q *Queries) SelectDefault(ctx context.Context) (pgtype.Text, error) {
	row := q.db.QueryRow(ctx, selectDefault)
	var default_ pgtype.Text
	err := row.Scan(&default_)
	return default_, err
}

const selectDefer = `-- name: SelectDefer :one
SELECT "defer" FROM go_keywords
`

func (q *Queries) SelectDefer(ctx context.Context) (pgtype.Text, error) {
	row := q.db.QueryRow(ctx, selectDefer)
	var defer_ pgtype.Text
	err := row.Scan(&defer_)
	return defer_, err
}

const selectElse = `-- name: SelectElse :one
SELECT "else" FROM go_keywords
`

func (q *Queries) SelectElse(ctx context.Context) (pgtype.Text, error) {
	row := q.db.QueryRow(ctx, selectElse)
	var else_ pgtype.Text
	err := row.Scan(&else_)
	return else_, err
}

const selectFallthrough = `-- name: SelectFallthrough :one
SELECT "fallthrough" FROM go_keywords
`

func (q *Queries) SelectFallthrough(ctx context.Context) (pgtype.Text, error) {
	row := q.db.QueryRow(ctx, selectFallthrough)
	var fallthrough_ pgtype.Text
	err := row.Scan(&fallthrough_)
	return fallthrough_, err
}

const selectFor = `-- name: SelectFor :one
SELECT "for" FROM go_keywords
`

func (q *Queries) SelectFor(ctx context.Context) (pgtype.Text, error) {
	row := q.db.QueryRow(ctx, selectFor)
	var for_ pgtype.Text
	err := row.Scan(&for_)
	return for_, err
}

const selectFunc = `-- name: SelectFunc :one
SELECT "func" FROM go_keywords
`

func (q *Queries) SelectFunc(ctx context.Context) (pgtype.Text, error) {
	row := q.db.QueryRow(ctx, selectFunc)
	var func_ pgtype.Text
	err := row.Scan(&func_)
	return func_, err
}

const selectGo = `-- name: SelectGo :one
SELECT "go" FROM go_keywords
`

func (q *Queries) SelectGo(ctx context.Context) (pgtype.Text, error) {
	row := q.db.QueryRow(ctx, selectGo)
	var go_ pgtype.Text
	err := row.Scan(&go_)
	return go_, err
}

const selectGoto = `-- name: SelectGoto :one
SELECT "goto" FROM go_keywords
`

func (q *Queries) SelectGoto(ctx context.Context) (pgtype.Text, error) {
	row := q.db.QueryRow(ctx, selectGoto)
	var goto_ pgtype.Text
	err := row.Scan(&goto_)
	return goto_, err
}

const selectIf = `-- name: SelectIf :one
SELECT "if" FROM go_keywords
`

func (q *Queries) SelectIf(ctx context.Context) (pgtype.Text, error) {
	row := q.db.QueryRow(ctx, selectIf)
	var if_ pgtype.Text
	err := row.Scan(&if_)
	return if_, err
}

const selectImport = `-- name: SelectImport :one
SELECT "import" FROM go_keywords
`

func (q *Queries) SelectImport(ctx context.Context) (pgtype.Text, error) {
	row := q.db.QueryRow(ctx, selectImport)
	var import_ pgtype.Text
	err := row.Scan(&import_)
	return import_, err
}

const selectInterface = `-- name: SelectInterface :one
SELECT "interface" FROM go_keywords
`

func (q *Queries) SelectInterface(ctx context.Context) (pgtype.Text, error) {
	row := q.db.QueryRow(ctx, selectInterface)
	var interface_ pgtype.Text
	err := row.Scan(&interface_)
	return interface_, err
}

const selectMap = `-- name: SelectMap :one
SELECT "map" FROM go_keywords
`

func (q *Queries) SelectMap(ctx context.Context) (pgtype.Text, error) {
	row := q.db.QueryRow(ctx, selectMap)
	var map_ pgtype.Text
	err := row.Scan(&map_)
	return map_, err
}

const selectPackage = `-- name: SelectPackage :one
SELECT "package" FROM go_keywords
`

func (q *Queries) SelectPackage(ctx context.Context) (pgtype.Text, error) {
	row := q.db.QueryRow(ctx, selectPackage)
	var package_ pgtype.Text
	err := row.Scan(&package_)
	return package_, err
}

const selectRange = `-- name: SelectRange :one
SELECT "range" FROM go_keywords
`

func (q *Queries) SelectRange(ctx context.Context) (pgtype.Text, error) {
	row := q.db.QueryRow(ctx, selectRange)
	var range_ pgtype.Text
	err := row.Scan(&range_)
	return range_, err
}

const selectReturn = `-- name: SelectReturn :one
SELECT "return" FROM go_keywords
`

func (q *Queries) SelectReturn(ctx context.Context) (pgtype.Text, error) {
	row := q.db.QueryRow(ctx, selectReturn)
	var return_ pgtype.Text
	err := row.Scan(&return_)
	return return_, err
}

const selectSelect = `-- name: SelectSelect :one
SELECT "select" FROM go_keywords
`

func (q *Queries) SelectSelect(ctx context.Context) (pgtype.Text, error) {
	row := q.db.QueryRow(ctx, selectSelect)
	var select_ pgtype.Text
	err := row.Scan(&select_)
	return select_, err
}

const selectStruct = `-- name: SelectStruct :one
SELECT "struct" FROM go_keywords
`

func (q *Queries) SelectStruct(ctx context.Context) (pgtype.Text, error) {
	row := q.db.QueryRow(ctx, selectStruct)
	var struct_ pgtype.Text
	err := row.Scan(&struct_)
	return struct_, err
}

const selectSwitch = `-- name: SelectSwitch :one
SELECT "switch" FROM go_keywords
`

func (q *Queries) SelectSwitch(ctx context.Context) (pgtype.Text, error) {
	row := q.db.QueryRow(ctx, selectSwitch)
	var switch_ pgtype.Text
	err := row.Scan(&switch_)
	return switch_, err
}

const selectType = `-- name: SelectType :one
SELECT "type" FROM go_keywords
`

func (q *Queries) SelectType(ctx context.Context) (pgtype.Text, error) {
	row := q.db.QueryRow(ctx, selectType)
	var type_ pgtype.Text
	err := row.Scan(&type_)
	return type_, err
}

const selectVar = `-- name: SelectVar :one
SELECT "var" FROM go_keywords
`

func (q *Queries) SelectVar(ctx context.Context) (pgtype.Text, error) {
	row := q.db.QueryRow(ctx, selectVar)
	var var_ pgtype.Text
	err := row.Scan(&var_)
	return var_, err
}
