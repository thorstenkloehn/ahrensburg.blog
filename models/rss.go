package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Rsstabele struct {
	ID  uint `gorm:"primaryKey"`
	Url string
}

func Start() {
	r := Rsstabele{}
	db, err := gorm.Open(sqlite.Open("Datenbank.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	r.Url = "https://ahrensburg-portal.de"
	db.AutoMigrate(&Rsstabele{})
	db.Create(&r)
}
