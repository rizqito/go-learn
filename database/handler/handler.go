package handler

import (
	"context"
	"fmt"
	"go-learn/database/config"
	"go-learn/database/model"
)

func InsertAnimal() {
	db := config.DatabaseInit()
	defer db.Close()

	ctx := context.Background()

	_, err := db.ExecContext(ctx, "INSERT INTO animal(name) VALUES('Anggora')")
	if err != nil {
		panic(err)
	}
	fmt.Println("Success Insert Data To Database")
}

func GetAnimal() {
	db := config.DatabaseInit()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT * FROM animal"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id : ", id)
		fmt.Println("Nama : ", name)
	}

	defer rows.Close()
}

func updateAnimal() {

}

func InsertUser() {
	db, err := config.GormDatabaseConn()

	if err != nil {
		fmt.Errorf("Error connect db", err.Error())
	}

	user := model.User{Name: "Sandy", Email: "sandy@yahoo.com", Address: "Houston"}

	db.Create(&user)

	fmt.Println("Insert User successfully")
}

func SelectUser() {
	db, err := config.GormDatabaseConn()

	if err != nil {
		fmt.Errorf("Error connect db", err.Error())
	}

	var users []model.User

	result := db.Find(&users)
}
