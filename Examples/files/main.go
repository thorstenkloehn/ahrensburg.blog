package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Website struct {
	Titel       []string
	Inhalt      string
	Verzeichnis []string
}

var vorlagen, _ = template.ParseGlob("Examples/files/views/*")

func main() {

	const home = "Examples/files/Datei"

	fmt.Println(filepath.Base(home))
	files, err := os.ReadDir(home)
	if err != nil {
		log.Fatal(err)
	}
	var Ausgabe []string
	var Verzeichnis []string
	for _, file := range files {

		if file.Type() == 0 {

			Ausgabe = append(Ausgabe, strings.Trim(file.Name(), ".md"))

		} else {
			Verzeichnis = append(Verzeichnis, file.Name())
		}

	}
	var website Website

	f, err := os.Create("datei/hallo.html")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	lesen, error := ioutil.ReadFile("Datei/index.md")
	if error != nil {
		log.Println(error)
	}
	website.Inhalt = string(lesen)
	website.Verzeichnis = []string(Verzeichnis)
	website.Titel = []string(Ausgabe)
	vorlagen.ExecuteTemplate(f, "index.html", &website)

}
