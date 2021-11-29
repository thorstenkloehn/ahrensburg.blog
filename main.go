package main

import (
	"fmt"
	"git.ahrensburg.digital/Thorsten/ahrensburg.blog/controller"
	"git.ahrensburg.digital/Thorsten/ahrensburg.blog/models"
	"git.ahrensburg.digital/Thorsten/ahrensburg.blog/module/Rechsliche_Angabe"
	"git.ahrensburg.digital/Thorsten/ahrensburg.blog/module/rss"
	"git.ahrensburg.digital/Thorsten/ahrensburg.blog/module/wordpress_pages"
	"github.com/robfig/cron"
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
	c.AddFunc("@hourly", func() { rss.Start() })
	c.Start()
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
