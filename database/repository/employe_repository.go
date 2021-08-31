package repository

import (
	"belajar-golang/database/model"
	"context"
)

type EmployeeRepository interface {
	Insert(ctx context.Context, employee model.Employee) (model.Employee, error)
	FindById(ctx context.Context, id int) (model.Employee, error)
	FindAll(ctx context.Context) ([]model.Employee, error)
}
