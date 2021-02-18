package structs

import "github.com/jinzhu/gorm"

type Slot struct {
	gorm.Model
	Availability int
	Name         string
}
