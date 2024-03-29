package sql

const (
	CreateParameterTable = `
	CREATE TABLE IF NOT EXISTS parameter (
		name text NOT NULL UNIQUE,
		value text
	)
	`

	CreateNgxStatusTable = `
	CREATE TABLE IF NOT EXISTS ngx_status (
		time datetime,
		connections int,
		requests int
	)
	`

	CreateNgxLogTable = `
	CREATE TABLE IF NOT EXISTS ngx_log (
		time datetime,
		json text
	)
	`
)
