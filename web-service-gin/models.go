package models

import "gorm.io/gorm"

// album represents data about a record album.
type Album struct {
	gorm.Model
	Title  string
	Artist string
	Price  float64
}
