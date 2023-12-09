package mysqlconn

import (
	"context"
	"fmt"
)

type MysqlService interface {
	CreateDatabase(ctx context.Context) (int64, error)
	CreateDatabaseWith(ctx context.Context, db string) (int64, error)
	ExecuteBatch(statements []string) error
	ExecuteBatchWithTransaction(statements []string) error
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

func (p *mysqlServiceImpl) ExecuteBatch(statements []string) error {
	if len(statements) == 0 {
		return fmt.Errorf("missing statements")
	}
	tx, err := p.mysqlConn.conn.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	for _, statement := range statements {
		_, err := tx.Exec(statement)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *mysqlServiceImpl) ExecuteBatchWithTransaction(statements []string) error {
	tx, err := p.mysqlConn.conn.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	err = p.ExecuteBatch(statements)
	if err != nil {
		return err
	}
	return nil
}
