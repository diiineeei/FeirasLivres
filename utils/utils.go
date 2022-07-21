package utils

import (
	"encoding/json"
	"net/http"
)

func JsonHttpResponse(w http.ResponseWriter, msg interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(msg)
}

func HttpResponseJson(w http.ResponseWriter, msg string, code int) {
	jmessage := JsonHttpResponseMessage{
		StatusCode: code,
		Message:    msg,
	}
	JsonHttpResponse(w, jmessage, code)
}

func String2Interface(s []string) []interface{} {
	i := make([]interface{}, len(s))
	for k, v := range s {
		i[k] = v
	}
	return i
}
