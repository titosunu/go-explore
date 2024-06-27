package repository

import "go-project/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}