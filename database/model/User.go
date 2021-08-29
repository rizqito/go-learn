package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID      uint
	Email   string
	Name    string
	Address string
}

type Product struct {
	gorm.Model
	ID       uint
	Name     string
	Qty      int32
	Supplier string
}
