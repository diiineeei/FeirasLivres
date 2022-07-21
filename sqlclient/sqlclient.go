package sqlclient

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

const (
	MaxSQLConnections = 50
	//defaultDataSource = "root:rodinei@tcp(mysql:3306)/rodinei"
	defaultDataSource = "root:rodinei@tcp(127.0.0.1:3306)/rodinei"
)

func Connect() (*sql.DB, error) {
	DataSource := os.Getenv("DATASOURCE")
	if DataSource == "" {
		DataSource = defaultDataSource
	}

	db, err := sql.Open("mysql", DataSource)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(MaxSQLConnections)

	return db, err
}
