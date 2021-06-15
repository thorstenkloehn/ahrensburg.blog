package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Website struct {
	gorm.Model
	Titel string
	Url   string
}

func main() {

	db, err := gorm.Open(sqlite.Open("Website.sql"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	var website []Website

	db.Find(&website)
	fmt.Println(website)

}
