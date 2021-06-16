package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	verzeichnis("./docs/01_Installieren")

}

func verzeichnis(Verzeichnis string) {
	files, _ := ioutil.ReadDir(Verzeichnis)
	for _, file := range files {
		if file.IsDir() == true {
			fmt.Println("Verzeichnis" + file.Name())
		} else {
			//	fmt.Println(file.Name())
		}

		fmt.Println(file.Name())

	}

}
