package partymix

import (
	"fmt"
	"net/http"
)

func Server() {
	fmt.Println("Le serveur est lancé : http://localhost:8080")
	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fscript := http.FileServer(http.Dir("./asset/"))
	http.Handle("/asset/", http.StripPrefix("/asset/", fscript))

	fp := http.FileServer(http.Dir("./pages/"))
	http.Handle("/pages/", http.StripPrefix("/pages/", fp))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Index(w, r)
	})

	http.ListenAndServe(":8080", nil)
}
