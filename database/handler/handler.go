package handler

import (
	"belajar-golang/database/config"
	"belajar-golang/database/model"
	"context"
	"fmt"
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

	fmt.Println("Insert user successfully")
}

func SelectUser() {
	db, err := config.GormDatabaseConn()

	if err != nil {
		fmt.Errorf("Error connect db", err.Error())
	}

	var user model.User

	result, err := db.Model(&model.User{}).Rows()

	if err != nil {
		fmt.Errorf("Cannot scan row", err.Error())
	}

	for result.Next() {
		db.ScanRows(result, &user)
		fmt.Println(user.Name)
		fmt.Println(user.Email)
	}

}

func UpdateUser() {
	db, err := config.GormDatabaseConn()

	if err != nil {
		fmt.Errorf("Error connect db", err.Error())
	}

	var user model.User

	db.Model(&user).Where("id = ?", 1).Update("name", "Aulia")

	fmt.Println("Update user sucess")
}

func DeleteUser() {
	db, err := config.GormDatabaseConn()

	if err != nil {
		fmt.Errorf("Error connect db", err.Error())
	}

	var user model.User

	db.Where("id = ?", 1).Delete(&user)

	fmt.Println("Delete user sucess")
}
