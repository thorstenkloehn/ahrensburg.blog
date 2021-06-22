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

		db, err := gorm.Open(sqlite.Open("Datenbank.db"), &gorm.Config{})
		if err != nil {

			log.Fatal(err)
		}
		urls := r.FormValue("urlSeiten")
		if urls == "" {
			http.Redirect(w, r, "http://localhost:8080", 301)
			return
		}
		daten, error1 := http.Get(urls)
		if error1 != nil {

		} else if daten.Status == "200 OK" {

			doc, err := htmlquery.LoadURL(urls)
			if err != nil {
				fmt.Fprintln(w, err)
			} else {
				a := htmlquery.FindOne(doc, "//title")
				fmt.Fprintln(w, htmlquery.InnerText(a))
				daten1, _ := http.Get(urls + "/wp-json/wp/v2/posts")
				if daten1.Status == "200 OK" {
					fmt.Fprintln(w, "Kann Eingeschrieben werden")
					wordpress := []models.Wordpress{{Url: urls, Titel: htmlquery.InnerText(a)}}
					db.Create(&wordpress)
				} else {
					fmt.Fprintln(w, "nicht eingeschriebnen werden")
				}
			}
		} else {
			fmt.Fprintln(w, "keine Daten")
		}
	}

}

func WordpressWebformulaloeshen(w http.ResponseWriter, r *http.Request) {
	if r.Host == "localhost:8080" {

	} else {
		fmt.Fprintln(w, "Sie haben kein Zugriff")
	}
}
