package handlers

import (
	"github.com/gorilla/mux"

	"testesrod/rest/feiraslivres"
	"testesrod/sqlclient"

	"net/http"
)

func New() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/importar", sqlclient.DataImport).Methods("GET")
	r.HandleFunc("/nova-feira", feiraslivres.InsereNovaFeira).Methods("POST")
	r.HandleFunc("/exclui-feira", feiraslivres.ExcluiFeira).Methods("POST")
	r.HandleFunc("/altera-feira", feiraslivres.AlteraFeira).Methods("POST")
	r.HandleFunc("/seleciona-feira", feiraslivres.SelecionaFeira).Methods("GET")
	return r
}
