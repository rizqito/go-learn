package repository

import "belajar-golang/testing/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
