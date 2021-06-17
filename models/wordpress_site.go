package models

type Wordpress struct {
	ID    uint `gorm:"primaryKey"`
	Url   string
	Titel string
}
