package mysqlconn

import "database/sql"

type MySql struct {
	conn *sql.DB `json:"-"`
}
