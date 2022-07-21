package sqlclient

import (
	_ "github.com/go-sql-driver/mysql"
	"testesrod/utils"

	"bytes"
	"database/sql"
	_ "embed"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"testesrod/logs"
	"time"
)

//go:embed feiraslivres.sql
var estruturaTabela []byte

//go:embed DEINFO_AB_FEIRASLIVRES_2014.csv
var dadosFeirasLivres []byte

const (
	TABLENAME = "FEIRASLIVRES"
	DELIMITER = ','
)

func CreateTable() error {
	db, err := Connect()
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec("DROP TABLE IF EXISTS " + TABLENAME)
	if err != nil {
		return err
	}
	_, err = db.Exec(string(estruturaTabela))
	return err
}

func DataImport(w http.ResponseWriter, r *http.Request) {
	logs.Print("Iniciando importação de dados")

	tempoCorrido, err := Import()
	if err != nil {
		utils.HttpResponseJson(w, err.Error(), 200)
		return
	}
	utils.HttpResponseJson(w, fmt.Sprintf("tempo usado para importar arquivo %s", tempoCorrido), 200)
}

func Import() (time.Duration, error) {
	logs.Print("Iniciando importação de dados")

	reader := csv.NewReader(bytes.NewReader(dadosFeirasLivres))
	reader.Comma = DELIMITER
	db, err := Connect()
	if err != nil {
		logs.Print("Erro ao tentar conectar ao Mysql", err)
		return 0, fmt.Errorf("erro ao tentar conectar ao Mysql")
	}

	defer db.Close()

	tempoInicial := time.Now()
	query := ""
	callback := make(chan int)
	connections := 0
	connDisponiveis := make(chan bool, MaxSQLConnections)
	for i := 0; i < MaxSQLConnections; i++ {
		connDisponiveis <- true
	}

	startConnectionController(&connections, callback, connDisponiveis)

	var wg sync.WaitGroup
	id := 1
	primeiraColuna := true

	for {
		var record []string
		record, err = reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			if len(record) == 16 {
				break
			}
			logs.Print("Erro ao importar dados do arquivo", err)
			return 0, fmt.Errorf("erro ao importar dados do arquivo")
		}
		if primeiraColuna {
			parseColumns(record, &query)
			primeiraColuna = false
		} else if <-connDisponiveis {
			connections += 1
			id += 1
			wg.Add(1)
			go insert(id, query, db, callback, &connections, &wg, utils.String2Interface(record))
		}
	}

	wg.Wait()

	tempoCorrido := time.Since(tempoInicial)
	logs.Print("tempo usado para importar arquivo", tempoCorrido)
	return tempoCorrido, nil
}

func insert(id int, query string, db *sql.DB, callback chan<- int, conns *int, wg *sync.WaitGroup, args []interface{}) {
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(args...)
	if err != nil {
		log.Printf("ID: %d (%d conns), %s\n", id, *conns, err.Error())
	}

	callback <- id
	wg.Done()
}

func startConnectionController(connections *int, callback <-chan int, available chan<- bool) {
	go func() {
		for {
			<-callback
			*connections -= 1
			available <- true
		}
	}()
}

func parseColumns(columns []string, query *string) {
	*query = "INSERT IGNORE INTO " + TABLENAME + " ("
	placeholder := "VALUES ("
	for i, c := range columns {
		if i == 0 {
			*query += c
			placeholder += "?"
		} else {
			*query += ", " + c
			placeholder += ", ?"
		}
	}
	placeholder += ")"
	*query += ") " + placeholder
}
