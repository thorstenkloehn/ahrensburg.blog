package Rechsliche_Angabe

import (
	"fmt"
	"github.com/russross/blackfriday/v2"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

type Rechtliche_Angabe struct {
	Titel  string
	Inhalt string
}

func Start() {
	var vorlagen, _ = template.ParseGlob("views/*")

	files, _ := ioutil.ReadDir("input/RechtlicheAngaben")
	rechtliche_Angabe := Rechtliche_Angabe{}
	for _, file := range files {

		lesen, err := ioutil.ReadFile("input/RechtlicheAngaben/" + file.Name())
		if err != nil {

			log.Println(err)
		}

		output := blackfriday.Run((lesen))

		f, err := os.Create("output/Rechtliche_Angabe/" + strings.TrimSuffix(file.Name(), ".md") + ".html")
		if err != nil {
			fmt.Println(err)
			f.Close()
			return
		}
		rechtliche_Angabe.Titel = strings.TrimSuffix(file.Name(), ".md")
		rechtliche_Angabe.Inhalt = string(output)

		vorlagen.ExecuteTemplate(f, "rechtliche_Angabe.html", &rechtliche_Angabe)
		f.Close()

	}

}
