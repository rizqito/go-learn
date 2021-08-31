package handler

import (
	"belajar-golang/database/config"
	"belajar-golang/database/model"
	"context"
	"fmt"
)

func InsertEmployee() {
	db := config.DatabaseInit()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO employee(name,address,age) VALUES('Riki','Jakarta,'65')"

	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success Insert Data To Database")
}

func GetEmployee() {
	db := config.DatabaseInit()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT * FROM employee"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name, address, age string
		err := rows.Scan(&id, &name, &address, &age)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id : ", id)
		fmt.Println("Nama : ", name)
		fmt.Println("Address : ", address)
		fmt.Println("Age : ", age)
	}

	defer rows.Close()
}

func UpdateEmployee() {
	db := config.DatabaseInit()
	defer db.Close()

	ctx := context.Background()

	script := "UPDATE employee SET name = 'Alex', address = 'Jarkata', age = '20', WHERE id = '2'"
	_, err := db.ExecContext(ctx, script)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success update new employee")
}

func DeleteEmployee() {
	db := config.DatabaseInit()
	defer db.Close()

	ctx := context.Background()

	id := 1

	script := "DELETE FROM employee WHERE id = ?"
	_, err := db.ExecContext(ctx, script, id)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success delete employee")
}

//* =====================GORM Handler=====================

func MigrateProduct() {
	db, err := config.GormDatabaseConn()

	if err != nil {
		fmt.Errorf("Error connect db", err.Error())
	}

	// db.Set("gorm:products", "ENGINE=InnoDB").AutoMigrate(&model.Product{})
	db.AutoMigrate(
		&model.Customer{},
		&model.Order{},
		&model.Shipment{},
	)
}

func InsertProduct() {
	db, err := config.GormDatabaseConn()

	if err != nil {
		fmt.Errorf("Error connect db", err.Error())
	}

	product := model.Product{Name: "Ultramilk", Qty: 25, Supplier: "PT. INdofood"}

	db.Create(&product)

	fmt.Println("Insert product successfully")
}

func SelectProduct() {
	db, err := config.GormDatabaseConn()

	if err != nil {
		fmt.Errorf("Error connect db", err.Error())
	}

	var product model.Product

	result, err := db.Model(&model.Product{}).Rows()

	if err != nil {
		fmt.Errorf("Cannot scan row", err.Error())
	}

	for result.Next() {
		db.ScanRows(result, &product)
		fmt.Println(product.Name)
		fmt.Println(product.Qty)
		fmt.Println(product.Supplier)
	}

}

func UpdateProduct() {
	db, err := config.GormDatabaseConn()

	if err != nil {
		fmt.Errorf("Error connect db", err.Error())
	}

	var product model.Product

	db.Model(&product).Where("id = ?", 1).Update("name", "Lifeboy")

	fmt.Println("Update product sucess")
}

func DeleteUser() {
	db, err := config.GormDatabaseConn()

	if err != nil {
		fmt.Errorf("Error connect db", err.Error())
	}

	var product model.Product

	db.Where("id = ?", 1).Delete(&product)

	fmt.Println("Delete product sucess")
}
