package example

import (
	"context"
	"testing"

	"github.com/sivaosorg/govm/dbx"
	"github.com/sivaosorg/govm/logger"
	"github.com/sivaosorg/govm/mysql"
	"github.com/sivaosorg/msqlconn"
)

func createConn() (*msqlconn.MySql, dbx.Dbx) {
	return msqlconn.NewClient(*mysql.GetMysqlConfigSample().SetDebugMode(true))
}

func TestConn(t *testing.T) {
	_, s := createConn()
	logger.Infof("Msql connection status: %v", s)
}

func TestCoreServiceCreateNewDatabase(t *testing.T) {
	m, _ := createConn()
	svc := msqlconn.NewMysqlService(m)
	_, err := svc.CreateDatabaseWith(context.Background(), "user")
	if err != nil {
		logger.Errorf("Creating database got an error", err)
		return
	}
}
