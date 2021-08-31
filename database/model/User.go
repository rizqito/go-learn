package model

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID       uint
	Name     string
	Qty      int32
	Supplier string
}

type Customer struct {
	gorm.Model
	Name  string
	Order []Order
}

type Order struct {
	gorm.Model
	CustomerId uint
	Shipment   []Shipment
	OrderDate  time.Time
}

type Shipment struct {
	gorm.Model
	OrderId      uint
	ShipmentDate time.Time
}

type Employee struct {
	Id      int
	Name    string
	Address string
	Age     int
}
