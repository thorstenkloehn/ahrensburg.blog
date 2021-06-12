package main

import (
	"fmt"
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
		test := Seite{Titel: file.Name(), Inhalt: "Hallo Leute"}
		ausgabe = append(ausgabe, test)
	}

	fmt.Println(ausgabe)
}
