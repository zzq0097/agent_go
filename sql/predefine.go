package sql

const (
	CreateParameterTable = `
	CREATE TABLE IF NOT EXISTS parameter (
		name text,
		value text
	)
	`

	CreateNgxStatusTable = `
	CREATE TABLE IF NOT EXISTS ngx_status (
		time text
		connections int,
		requests int
	)
	`

	CreateNgxCodeTable = `
	CREATE TABLE IF NOT EXISTS ngx_code (
		time text,
		json text
	)
	`
)
