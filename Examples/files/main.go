package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

type Website struct {
	Titel  []string
	Inhalt string
}

var vorlagen, _ = template.ParseGlob("Examples/files/views/*")

func main() {

	files, err := os.ReadDir("datei")
	if err != nil {
		log.Fatal(err)
	}
	var Ausgabe []string
	for _, file := range files {

		Ausgabe = append(Ausgabe, file.Name())

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
	website.Titel = []string(Ausgabe)

	vorlagen.ExecuteTemplate(f, "index.html", &website)

}
