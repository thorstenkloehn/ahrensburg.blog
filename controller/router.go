package controller

import "net/http"

type Seite struct {
	Titel  string
	Inhalt string
}

func Login(w http.ResponseWriter, r *http.Request) {

	title := r.URL.Path[len("/login/"):]
	w.Write([]byte(title))
	if r.Method == "localhost" {
		w.Write([]byte("Sie haben kein Zugriff auf der Datenbsnk"))
	} else {
		w.Write([]byte("Sie haben Zugriff auf der Datenbank"))
	}
}
