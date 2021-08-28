package config

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DatabaseInit() *sql.DB {
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

func GormDatabaseConn() (*gorm.DB, error) {
	user := "root"
	pass := ""
	port := "3306"
	host := "127.0.0.1"
	dbName := "go_db"

	dsn := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Errorf("Cannot connect db", err)
		panic(err)
	}

	return db, nil
}
