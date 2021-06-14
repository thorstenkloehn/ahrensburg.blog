package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Webseite struct {
	Titel       string
	Menue []string
	Inhalt      []string
	Verzeichnis []string
}

var vorlagen, _ = template.ParseGlob("Examples/wordpress_pages/views/*")

func main() {
	const home = "Examples/files/datei"
	const ausgabe = "Examples/files/ausgabe"
    homepathdatei := filepath.Join(home)
    // homepathausgabe : filepath.Join(ausgabe)
	files, err := os.ReadDir(homepathdatei)
	if err != nil {
		log.Fatal(err)
	}
	var Ausgabe []string
	var Verzeichnis [] string
	for _, file := range files {

		if file.Type() == 0 {
			Ausgabe = append(Ausgabe, strings.Trim(file.Name(), ".md"))
		} else {
			Verzeichnis = append(Verzeichnis,file.Name())

		}


	}

	// fmt.Println(Ausgabe)
	var webseitem =Webseite{}
	for i,_:=range Ausgabe {
		website.Titel = Ausgabe[i]
		website.Menue = Ausgabe

	}

	f, err := os.Create("wordpress_pages/out/index.html")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	vorlagen.ExecuteTemplate(f, "index.html", &webseitem)
	f.Close()



	}




