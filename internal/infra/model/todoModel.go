package model

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Task   string
	Status bool
}
