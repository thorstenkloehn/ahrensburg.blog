package main

import (
	"fmt"
	"github.com/thorstenkloehn/ahrensburg.blog/controller"
	"github.com/thorstenkloehn/ahrensburg.blog/models"
	"github.com/thorstenkloehn/ahrensburg.blog/module/Rechsliche_Angabe"
	"github.com/thorstenkloehn/ahrensburg.blog/module/rss"
	"github.com/thorstenkloehn/ahrensburg.blog/module/wordpress_pages"
	"net/http"
	"os"
)

func main() {

	_, errordatei := os.Stat("Datenbank.db")
	if errordatei != nil {
		if os.IsNotExist(errordatei) {
			wordpress := models.Wordpress{}
			wordpress.NeueTabelleerstellen()

		}
	}

	// db.AutoMigrate(&models.Wordpress{})

	wordpress_pages.Start()
	Rechsliche_Angabe.Start()
	rss.Start()

	router := http.NewServeMux()
	router.HandleFunc("/formular", controller.WordpressWebformular)
	router.HandleFunc("/loeschen", controller.WordpressWebformulaloeshen)
	router.HandleFunc("/", controller.Login)
	fmt.Println(" http://localhost:8080")
	http.ListenAndServe("localhost:8080", router)
	//Fertige
}
