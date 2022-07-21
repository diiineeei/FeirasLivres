package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func New() http.Handler {
	r := mux.NewRouter()

	return r
}
