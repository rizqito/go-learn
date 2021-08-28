package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID      uint
	Email   string
	Name    string
	Address string
}
