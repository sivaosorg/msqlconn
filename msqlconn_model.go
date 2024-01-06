package msqlconn

import (
	"database/sql"

	"github.com/sivaosorg/govm/dbx"
	"github.com/sivaosorg/govm/mysql"
)

type MySql struct {
	conn   *sql.DB           `json:"-"`
	Config mysql.MysqlConfig `json:"config,omitempty"`
	State  dbx.Dbx           `json:"state,omitempty"`
}
