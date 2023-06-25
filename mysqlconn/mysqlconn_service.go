package mysqlconn

import (
	"context"
	"fmt"
)

type MysqlService interface {
	CreateDatabase(ctx context.Context) (int64, error)
	CreateDatabaseWith(ctx context.Context, db string) (int64, error)
}

type mysqlServiceImpl struct {
	mysqlConn *MySql
}

func NewMysqlService(mysqlConn *MySql) MysqlService {
	s := &mysqlServiceImpl{
		mysqlConn: mysqlConn,
	}
	return s
}

func (m *mysqlServiceImpl) CreateDatabase(ctx context.Context) (int64, error) {
	return m.CreateDatabaseWith(ctx, m.mysqlConn.Config.Database)
}

func (m *mysqlServiceImpl) CreateDatabaseWith(ctx context.Context, db string) (int64, error) {
	if !m.mysqlConn.State.IsConnected {
		return -1, m.mysqlConn.State.Error
	}
	response, err := m.mysqlConn.conn.ExecContext(ctx, fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", db))
	if err != nil {
		return -1, err
	}
	no, err := response.RowsAffected()
	return no, err
}
