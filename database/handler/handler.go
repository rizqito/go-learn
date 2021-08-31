package handler

import (
	"belajar-golang/database/config"
	"belajar-golang/database/model"
	"context"
	"fmt"
	"strconv"
)

func LoginUser() {
	db := config.DatabaseInit()
	defer db.Close()

	ctx := context.Background()

	username := "admin"
	password := "admin"

	query := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, query, username, password)
	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("sukses login ", username)
	} else {
		fmt.Println("Gagal login")
	}

	defer rows.Close()
}

func LastInsertId() {
	db := config.DatabaseInit()
	defer db.Close()

	ctx := context.Background()

	name := "Onii"
	address := "Semarang"
	age := 70

	query := "INSERT INTO employee(name,address,age) VALUES(?,?,?)"

	result, err := db.ExecContext(ctx, query, name, address, age)
	if err != nil {
		panic(err)
	}
	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("success get last data id", insertId)
}

// input query berkali2 cocok pake prepare
func InsertPrepare() {
	db := config.DatabaseInit()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO employee(name,address,age) VALUES(?,?,?)"

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for i := 0; i < 10; i++ {
		name := "alex" + strconv.Itoa(i) + "@gmail.com"
		address := "address" + strconv.Itoa(i)
		age := i
		result, err := stmt.ExecContext(ctx, name, address, age)
		if err != nil {
			panic(err)
		}
		lastInsertId, _ := result.LastInsertId()
		fmt.Println("last id :", lastInsertId)
	}
}

func InsertEmployee() {
	db := config.DatabaseInit()
	defer db.Close()

	ctx := context.Background()

	name := "Riki"
	address := "Jakarta"
	age := 65

	query := "INSERT INTO employee(name,address,age) VALUES(?,?,?)"

	_, err := db.ExecContext(ctx, query, name, address, age)
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

	id := 2
	name := "Alex"
	address := "Jakarta"
	age := 20

	script := "UPDATE employee SET name = ?, address = ?, age = ?, WHERE id = ?"
	_, err := db.ExecContext(ctx, script, name, address, age, id)

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
