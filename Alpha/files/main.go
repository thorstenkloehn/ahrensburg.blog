package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Verzeichnis struct {
	Verzeichnis string
}

type Webseiten struct {
	Verzeichnis Verzeichnis
	Titel       string
	Inhalt      string
	Menue       []string
}

type Website struct {
	Verzeichnisse []Verzeichnis
	Webseiten     []Webseiten
	Titel         string
}

func main() {
	const home = "Examples/files/datei/"
	homepathdatei := filepath.Join(home)
	const ausgabe = "Examples/files/ausgabe"

	files, err := os.ReadDir(homepathdatei)
	if err != nil {
		log.Fatal(err)
	}
	webseiten := []Webseiten{}
	webseiten1 := Webseiten{}
	website := Website{}
	for _, file := range files {
		webseiten1.Titel = file.Name()
		if file.Type() == 0 {
			inhalt, _ := ioutil.ReadFile(homepathdatei + "/" + file.Name())
			webseiten1.Verzeichnis.Verzeichnis = file.Name()
			webseiten1.Menue = append(webseiten1.Menue, strings.Trim(file.Name(), ".md"))
			webseiten1.Inhalt = string(inhalt)
		} else {
			webseiten1.Verzeichnis.Verzeichnis = file.Name()
			webseiten1.Inhalt = ""
			webseiten1.Menue = []string{}
			webseiten1.Verzeichnis.Verzeichnis = ""
			website.Verzeichnisse = append(website.Verzeichnisse, webseiten1.Verzeichnis)

		}

		webseiten = append(webseiten, webseiten1)
	}
	website.Webseiten = webseiten
	for i, ttt := range website.Webseiten {
		fmt.Println(i, ttt.Titel, ttt.Menue, ttt.Inhalt)
	}
}
