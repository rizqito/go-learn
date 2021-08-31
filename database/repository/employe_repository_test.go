package repository

import (
	"belajar-golang/database/config"
	"belajar-golang/database/model"
	"context"
	"fmt"
	"testing"
)

func TestEmployeeInsert(t *testing.T) {
	employeeRepository := NewEmployeeRepository(config.DatabaseInit())
	ctx := context.Background()
	employee := model.Employee{
		Name:    "JOnas",
		Address: "JL. manggis",
		Age:     23,
	}
	result, err := employeeRepository.Insert(ctx, employee)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	employeeRepository := NewEmployeeRepository(config.DatabaseInit())
	employee, err := employeeRepository.FindById(context.Background(), 5)
	if err != nil {
		panic(err)
	}
	fmt.Println(employee)
}

func TestFindAll(t *testing.T) {
	employeeRepository := NewEmployeeRepository(config.DatabaseInit())
	employes, err := employeeRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, employee := range employes {
		fmt.Println(employee)
	}
}
