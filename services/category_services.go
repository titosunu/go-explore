package services

import (
	"errors"
	"go-project/entity"
	"go-project/repository"
)

type CategoryServices struct {
	Repository repository.CategoryRepository
}

func(service CategoryServices) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return category, errors.New("category not found bolo")
	} else {
		return category, nil
	}
}