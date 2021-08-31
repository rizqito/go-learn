package repository

import (
	"belajar-golang/database/model"
	"context"
	"database/sql"
	"errors"
	"strconv"
)

type employeeRepositoryImpl struct {
	DB *sql.DB
}

func NewEmployeeRepository(db *sql.DB) EmployeeRepository {
	return &employeeRepositoryImpl{DB: db}
}

func (repository *employeeRepositoryImpl) Insert(ctx context.Context, employee model.Employee) (model.Employee, error) {
	query := "INSERT INTO employee(name,address,age) VALUES(?,?,?)"
	result, err := repository.DB.ExecContext(ctx, query, employee.Name, employee.Address, employee.Age)
	if err != nil {
		return employee, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return employee, err
	}
	employee.Id = int(id)
	return employee, nil
}

func (repository *employeeRepositoryImpl) FindById(ctx context.Context, id int) (model.Employee, error) {
	query := "SELECT id, name, address, age FROM employee WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, query, id)
	employee := model.Employee{}
	if err != nil {
		return employee, err
	}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&employee.Id, &employee.Name, &employee.Address, &employee.Age)
		return employee, nil
	} else {
		return employee, errors.New("id" + strconv.Itoa(id) + "tidak ditemukan")
	}
}

func (repository *employeeRepositoryImpl) FindAll(ctx context.Context) ([]model.Employee, error) {
	query := "SELECT id, name, address, age FROM employee"
	rows, err := repository.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var employes []model.Employee
	for rows.Next() {
		employee := model.Employee{}
		rows.Scan(&employee.Id, &employee.Name, &employee.Address, &employee.Age)
		employes = append(employes, employee)
	}
	return employes, nil
}
