package controller

import (
	"fmt"
	"github.com/thorstenkloehn/ahrensburg.digital/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
	"text/template"
)

var vorlagen, _ = template.ParseGlob("views/*")

func Login(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, r.Host)
	// title := r.URL.Path[len("/login/")

	if r.Host == "localhost:8080" {
		wordpress := []models.Wordpress{}
		db, err := gorm.Open(sqlite.Open("Datenbank.db"), &gorm.Config{})
		if err != nil {

			log.Fatal(err)
		}
		db.Find(&wordpress).Row()
		fmt.Println(wordpress)
		vorlagen.ExecuteTemplate(w, "wordpressformularEintrag.html", wordpress)
	} else {
		fmt.Fprint(w, "Sie heben kein Zugriff")
	}

}
