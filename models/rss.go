package models

type Rsstabele struct {
	ID    uint `gorm:"primaryKey"`
	Url   string
	Titel string
}
