package partymix

import (
	"net/http"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./pages/index.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}
}
