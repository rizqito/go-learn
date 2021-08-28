package main

import (
	"go-learn/database/config"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// handler.InsertAnimal()
	// handler.GetAnimal()

	db, err := config.GormDatabaseConn()
}
