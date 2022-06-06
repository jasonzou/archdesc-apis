package sqlite

import (
	// imports the driver, don't remove this comment, golint requires.
	_ "github.com/mattn/go-sqlite3"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const sqlite3DriverName = "sqlite3"

// New returns a postgres connection.
func New(datasource string, opts ...sqlx.SqlOption) sqlx.SqlConn {
	return sqlx.NewSqlConn(sqlite3DriverName, datasource, opts...)
}
