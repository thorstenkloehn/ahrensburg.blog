package models

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Wordpress struct {
	ID    uint `gorm:"primaryKey"`
	Url   string
	Titel string
}

func (w *Wordpress) NeueTabelleerstellen() {
	db, err := gorm.Open(sqlite.Open("Datenbank.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	w.Titel = "Ahrensburg Portal"
	w.Url = "https://ahrensburg-portal.de"
	fmt.Println("Neue Datenbank wird erstellt")
	db.AutoMigrate(&Wordpress{})
	db.Create(&w)

}
