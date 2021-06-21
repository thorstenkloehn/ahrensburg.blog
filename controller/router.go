package controller

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/thorstenkloehn/ahrensburg.digital/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
	"text/template"
)

var vorlagen, _ = template.ParseGlob("views/*")

func Login(w http.ResponseWriter, r *http.Request) {

	// fmt.Fprint(w, r.Host)
	// title := r.URL.Path[len("/login/")

	if r.Host == "localhost:8080" {
		wordpress := []models.Wordpress{}
		db, err := gorm.Open(sqlite.Open("Datenbank.db"), &gorm.Config{})
		if err != nil {

			log.Fatal(err)
		}
		db.Find(&wordpress).Row()
		vorlagen.ExecuteTemplate(w, "wordpressformularEintrag.html", wordpress)
	} else {
		fmt.Fprint(w, "Sie heben kein Zugriff")
	}

}

func WordpressWebformular(w http.ResponseWriter, r *http.Request) {
	if r.Host == "localhost:8080" {
		/*	wordpress := []models.Wordpress{}
			db, err := gorm.Open(sqlite.Open("Datenbank.db"), &gorm.Config{})
			if err != nil {

				log.Fatal(err)
			}
		*/
		urls := r.FormValue("urlSeiten")
		if urls == "" {
			fmt.Println("kein Url")
			http.Redirect(w, r, "http:/localhost:8080", 301)

		} else {
			doc, err := htmlquery.LoadURL(urls)
			if err != nil {
				log.Println(err)
			}
			a := htmlquery.FindOne(doc, "//title")
			if a != nil {
				fmt.Println(htmlquery.InnerText(a))
			} else {

				fmt.Println("Fehler")

			}
		}

	}
}
