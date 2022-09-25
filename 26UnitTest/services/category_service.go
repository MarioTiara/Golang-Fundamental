package services

import (
	"errors"
	"unittest/entity"
	"unittest/repository"
)

type CategoryService struct {
	Respository repository.CategoryRepository
}

func (service *CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Respository.FindById(id)
	if category == nil {
		return category, errors.New("Category Not Found")
	} else {
		return category, nil
	}
}
