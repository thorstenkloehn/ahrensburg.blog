package controller

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/thorstenkloehn/ahrensburg.digital/models"
	"github.com/thorstenkloehn/ahrensburg.digital/module/feed"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
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
	//Zugriff Kontrolle
	if r.Host == "localhost:8080" {
		// Sie haben Zugriff
		urls := r.FormValue("urlSeiten")
		// fmt.Fprintln(w, "Sie haben Zugriff")
		//String Leer Formular
		if urls == "" {
			//	fmt.Fprintln(w, "Formular Leer")
		} else {

			//	fmt.Fprintln(w, "formular wahr ein Text")
			daten, error1 := http.Get(urls)
			if error1 != nil {
				//		fmt.Fprint(w, "wahr kein Urls")
			} else {
				if daten.Status == "200 OK" {
					doc, err := htmlquery.LoadURL(urls)
					if err != nil {
						//			fmt.Fprintln(w, err)
					} else {
						//		a := htmlquery.FindOne(doc, "//title")
						//		fmt.Fprintln(w, htmlquery.InnerText(a))
					}
					///weiter
					daten1, _ := http.Get(urls + "/wp-json/wp/v2/posts")
					if daten1.Status == "200 OK" {
						//	fmt.Fprintln(w, "Kann Eingeschrieben werden")
						db, err := gorm.Open(sqlite.Open("Datenbank.db"), &gorm.Config{})
						if err != nil {

							log.Fatal(err)
						}
						a := htmlquery.FindOne(doc, "//title")

						wordpress := []models.Wordpress{{Url: urls, Titel: htmlquery.InnerText(a)}}
						db.Create(&wordpress)

					} else {
						//					fmt.Fprintln(w, "Ist kein Wordüress Seite")
					}
				}
			}

		}
		http.Redirect(w, r, "http://localhost:8080", http.StatusFound)

	} else {
		//Sie habem kein Zugriff
		fmt.Fprint(w, "Sie haben kein Zugriff")
	}

}

func WordpressWebformulaloeshen(w http.ResponseWriter, r *http.Request) {
	if r.Host == "localhost:8080" {

		db, err := gorm.Open(sqlite.Open("Datenbank.db"), &gorm.Config{})
		if err != nil {

			log.Fatal(err)
		}
		ergebnis, _ := strconv.Atoi(r.FormValue(("Auswahl")))
		db.Delete(&models.Wordpress{}, ergebnis)
		http.Redirect(w, r, "http://localhost:8080", http.StatusFound)
	} else {
		fmt.Fprintln(w, "Sie haben kein Zugriff")
		http.Redirect(w, r, "http://localhost:8080", http.StatusFound)
	}

}

func FeedSeiten(w http.ResponseWriter, r *http.Request) {
	start := feed.Rss{}
	start.Lesen("http://ahrensburg-blog.de/feed")
	titel := start.Channel.Title
	fmt.Fprintln(w, titel)
}
