package handlers

import (
	"github.com/diiineeei/FeirasLivres/doc"
	"github.com/diiineeei/FeirasLivres/rest/feiraslivres"
	"github.com/diiineeei/FeirasLivres/sqlclient"
	"github.com/gorilla/mux"

	"net/http"
)

func New() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/", doc.Index).Methods("GET")
	r.HandleFunc("/importar", sqlclient.DataImport).Methods("GET")
	r.HandleFunc("/nova-feira", feiraslivres.InsereNovaFeira).Methods("POST")
	r.HandleFunc("/exclui-feira", feiraslivres.ExcluiFeira).Methods("POST")
	r.HandleFunc("/altera-feira", feiraslivres.AlteraFeira).Methods("POST")
	r.HandleFunc("/seleciona-feira", feiraslivres.SelecionaFeira).Methods("GET")
	return r
}
