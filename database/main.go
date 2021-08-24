package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func getConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/go_db")

	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(10)                  // pengaturan berapa minimal jumlah koneksi yang dibuat
	db.SetMaxOpenConns(100)                 // pengaturan berapa jumlah koneksi maksimal yang dibuat
	db.SetConnMaxIdleTime(5 * time.Minute)  // pengaturan berapa lama koneksi yang sudah tidak digunakan akan dihapus
	db.SetConnMaxLifetime(60 * time.Minute) // pengaturan berapa lama koneksi boleh digunakan

	return db
}

func createData() {
	db := getConnection()
	defer db.Close()

	ctx := context.Background()

	_, err := db.ExecContext(ctx, "INSERT INTO customer(id, name) VALUES('2','Joana');")
	if err != nil {
		panic(err)
	}
	fmt.Println("Success Insert Data To Database")
}

func viewData() {
	db := getConnection()
	defer db.Close()

	ctx := context.Background()

	rows, err := db.QueryContext(ctx, "SELECT * FROM customer")
	if err != nil {
		panic(err)
	}

	defer rows.Close()
}

func main() {
	// createData()
	viewData()
}
