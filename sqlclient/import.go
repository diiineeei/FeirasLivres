package sqlclient

import (
	_ "embed"
	_ "github.com/go-sql-driver/mysql"
)

//go:embed feiraslivres.sql
var estruturaTabela []byte

func CreateTable() error {
	db, err := Connect()
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec(string(estruturaTabela))
	return err
}
