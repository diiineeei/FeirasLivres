package feiraslivres

import (
	"github.com/diiineeei/FeirasLivres/logs"
	"github.com/diiineeei/FeirasLivres/sqlclient"
	"github.com/diiineeei/FeirasLivres/utils"

	"bytes"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func SelecionaFeira(w http.ResponseWriter, r *http.Request) {
	logs.Print("Iniciando importação de dados")

	distrito := r.Form.Get("distrito")
	regiao5 := r.Form.Get("regiao5")
	nomeFeira := r.Form.Get("nome_feira")
	bairro := r.Form.Get("bairro")

	feiras, err := SelectFeira(distrito, regiao5, nomeFeira, bairro)
	if err != nil {
		logs.Print("Erro ao atualizar Feira", err)
		utils.HttpResponseJson(w, "Erro ao atualizar Feira", 500)
		return
	}

	utils.JsonHttpResponse(w, feiras, 200)
}

func AlteraFeira(w http.ResponseWriter, r *http.Request) {
	logs.Print("Iniciando importação de dados")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logs.Print("Erro ao decodificar requisição", err)
		utils.HttpResponseJson(w, "Erro ao decodificar requisição", 400)
		return
	}

	var feiraLivreUpdate FeiraLivre
	err = json.Unmarshal(body, &feiraLivreUpdate)
	if err != nil {
		logs.Print("Erro ao decodificar JSON", err)
		utils.HttpResponseJson(w, "Erro ao decodificar JSON", 400)
		return
	}

	err = UpdateFeira(feiraLivreUpdate)
	if err != nil {
		logs.Print("Erro ao atualizar Feira", err)
		utils.HttpResponseJson(w, "Erro ao atualizar Feira", 500)
		return
	}

	utils.HttpResponseJson(w, "Feira atualizada com sucesso!", 200)
}

func ExcluiFeira(w http.ResponseWriter, r *http.Request) {
	logs.Print("Iniciando importação de dados")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logs.Print("Erro ao decodificar requisição", err)
		utils.HttpResponseJson(w, "Erro ao decodificar requisição", 400)
		return
	}

	var idFeiraLivre IDFeiraLivre
	err = json.Unmarshal(body, &idFeiraLivre)
	if err != nil {
		logs.Print("Erro ao decodificar JSON", err)
		utils.HttpResponseJson(w, "Erro ao decodificar JSON", 400)
		return
	}

	err = DeleteFeira(idFeiraLivre)
	if err != nil {
		logs.Print("Erro ao Deletar Feira", err)
		utils.HttpResponseJson(w, "Erro ao Deletar Feira", 500)
		return
	}

	utils.HttpResponseJson(w, "Feira Excluida com sucesso!", 200)
}

func InsereNovaFeira(w http.ResponseWriter, r *http.Request) {
	logs.Print("Iniciando importação de dados")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logs.Print("Erro ao decodificar requisição", err)
		utils.HttpResponseJson(w, "Erro ao decodificar requisição", 400)
		return
	}

	var feiraLivre FeiraLivreInsert
	err = json.Unmarshal(body, &feiraLivre)
	if err != nil {
		logs.Print("Erro ao decodificar JSON", err)
		utils.HttpResponseJson(w, "Erro ao decodificar JSON", 400)
		return
	}

	err = InsertFeira(feiraLivre)
	if err != nil {
		logs.Print("Erro ao Inserir nova Feira", err)
		utils.HttpResponseJson(w, "Erro ao Inserir nova Feira", 500)
		return
	}

	utils.HttpResponseJson(w, "Feira cadastrada com sucesso!", 200)
}

func InsertFeira(args FeiraLivreInsert) error {
	db, err := sqlclient.Connect()
	query := "INSERT INTO FEIRASLIVRES (LONGITUDE, LATITUDE, SETCENS, AREAP, CODDIST, DISTRITO, CODSUBPREF, SUBPREFE, REGIAO5, REGIAO8, NOME_FEIRA, REGISTRO, LOGRADOURO, NUMERO, BAIRRO, REFERENCIA) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

	stmt, err := db.Prepare(query)
	if err != nil {
		logs.Print("Erro ao executar db.Prepare", err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(args.LONGITUDE, args.LATITUDE, args.SETCENS, args.AREAP, args.CODDIST, args.DISTRITO, args.CODSUBPREF, args.SUBPREFE, args.REGIAO5, args.REGIAO8, args.NOME_FEIRA, args.REGISTRO, args.LOGRADOURO, args.NUMERO, args.BAIRRO, args.REFERENCIA)
	if err != nil {
		logs.Print("Erro ao executar insert ", err.Error())
		return err
	}
	return nil
}

func DeleteFeira(idFeira IDFeiraLivre) error {
	db, err := sqlclient.Connect()
	query := "DELETE FROM FEIRASLIVRES WHERE ID = ?"

	stmt, err := db.Prepare(query)
	if err != nil {
		logs.Print("Erro ao executar db.Prepare", err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(idFeira.ID)
	if err != nil {
		logs.Print("Erro ao executar insert ", err.Error())
		return err
	}

	return nil
}

func UpdateFeira(args FeiraLivre) error {
	db, err := sqlclient.Connect()
	query := "UPDATE FEIRASLIVRES SET LONGITUDE = ?, LATITUDE = ?, SETCENS = ?, AREAP = ?, CODDIST = ?, DISTRITO = ?, CODSUBPREF = ?, SUBPREFE = ?, REGIAO5 = ?, REGIAO8 = ?, NOME_FEIRA = ?, REGISTRO = ?, LOGRADOURO = ?, NUMERO = ?, BAIRRO = ?, REFERENCIA = ? WHERE ID = ?"

	stmt, err := db.Prepare(query)
	if err != nil {
		logs.Print("Erro ao executar db.Prepare", err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(args.LONGITUDE, args.LATITUDE, args.SETCENS, args.AREAP, args.CODDIST, args.DISTRITO, args.CODSUBPREF, args.SUBPREFE, args.REGIAO5, args.REGIAO8, args.NOME_FEIRA, args.REGISTRO, args.LOGRADOURO, args.NUMERO, args.BAIRRO, args.REFERENCIA, args.ID)
	if err != nil {
		logs.Print("Erro ao executar insert ", err.Error())
		return err
	}

	return nil
}

func SelectFeira(paramDistrito, paramRegiao5, paramNomeFeira, paramBairro string) ([]FeiraLivre, error) {
	db, err := sqlclient.Connect()
	if err != nil {
		return nil, err
	}

	var (
		query  bytes.Buffer
		params []interface{}
		feiras []FeiraLivre
	)

	query.WriteString("SELECT ID, LONGITUDE, LATITUDE, SETCENS, AREAP, CODDIST, DISTRITO, CODSUBPREF, SUBPREFE, REGIAO5, REGIAO8, NOME_FEIRA, REGISTRO, LOGRADOURO, NUMERO, BAIRRO, REFERENCIA FROM FEIRASLIVRES WHERE 1=1 ")
	if paramDistrito != "" {
		query.WriteString(" AND DISTRITO = ? ")
		params = append(params, paramDistrito)
	}
	if paramRegiao5 != "" {
		query.WriteString(" AND REGIAO5 = ? ")
		params = append(params, paramRegiao5)
	}
	if paramNomeFeira != "" {
		query.WriteString(" AND NOME_FEIRA = ? ")
		params = append(params, paramNomeFeira)
	}
	if paramBairro != "" {
		query.WriteString(" AND BAIRRO = ? ")
		params = append(params, paramBairro)
	}

	rows, err := db.Query(query.String(), params...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var (
			id, longitude, latitude, setcens, areap, coddist, codsubpref                                      sql.NullInt64
			distrito, subprefe, regiao5, regiao8, nomeFeira, registro, logradouro, numero, bairro, referencia sql.NullString
		)

		if err := rows.Scan(&id, &longitude, &latitude, &setcens, &areap, &coddist, &distrito, &codsubpref, &subprefe, &regiao5, &regiao8, &nomeFeira, &registro, &logradouro, &numero, &bairro, &referencia); err != nil {
			continue
		}
		feiras = append(feiras, FeiraLivre{
			ID:         int(id.Int64),
			LONGITUDE:  int(longitude.Int64),
			LATITUDE:   int(latitude.Int64),
			SETCENS:    int(setcens.Int64),
			AREAP:      int(areap.Int64),
			CODDIST:    int(coddist.Int64),
			DISTRITO:   distrito.String,
			CODSUBPREF: int(codsubpref.Int64),
			SUBPREFE:   subprefe.String,
			REGIAO5:    regiao5.String,
			REGIAO8:    regiao8.String,
			NOME_FEIRA: nomeFeira.String,
			REGISTRO:   registro.String,
			LOGRADOURO: logradouro.String,
			NUMERO:     numero.String,
			BAIRRO:     bairro.String,
			REFERENCIA: referencia.String,
		})
	}

	return feiras, nil
}
