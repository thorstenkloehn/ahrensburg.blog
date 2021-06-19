package main

import (
	"github.com/robfig/cron/v3"
	"github.com/thorstenkloehn/ahrensburg.digital/controller"
	"github.com/thorstenkloehn/ahrensburg.digital/models"
	"github.com/thorstenkloehn/ahrensburg.digital/module/Rechsliche_Angabe"
	"github.com/thorstenkloehn/ahrensburg.digital/module/wordpress_pages"
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
	c := cron.New()
	c.AddFunc("@hourly", func() { wordpress_pages.Start() })
	c.Start()
	wordpress_pages.Start()
	Rechsliche_Angabe.Start()

	router := http.NewServeMux()
	// router.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("output/wordpress"))))
	// router.Handle("/rechtliche_Angabe/", http.StripPrefix("/rechtliche_Angabe/", http.FileServer(http.Dir("output/Rechtliche_Angabe"))))
	router.HandleFunc("/login/", controller.Login)
	http.ListenAndServe("localhost:8080", router)
}
