package controller

import "net/http"

type Seite struct {
	Titel  string
	Inhalt string
}

func Login(w http.ResponseWriter, r *http.Request) {

	// title := r.URL.Path[len("/login/")
	if r.Host == "localhost" {
		w.Write([]byte("Sie haben  Zugriff auf die Datenbank"))

	} else {
		w.Write([]byte("Sie haben kein Zugriff auf die Datenbank"))
	}
}
