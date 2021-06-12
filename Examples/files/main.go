package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Seite struct {
	Titel  string
	Inhalt string
}

func main() {

	files, err := os.ReadDir("datei")
	if err != nil {
		log.Fatal(err)
	}
	ausgabe := []Seite{}

	for _, file := range files {

		inhaltContent, _ := ioutil.ReadFile("Datei/" + file.Name())
		inhalt := Seite{Titel: file.Name(), Inhalt: string(inhaltContent)}
		ausgabe = append(ausgabe, inhalt)
	}

	fmt.Println(ausgabe)
}
