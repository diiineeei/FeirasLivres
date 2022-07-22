package doc

import (
	_ "embed"
	"net/http"
)

//go:embed doc.html
var staticDocHtml []byte

func Index(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write(staticDocHtml)
}
