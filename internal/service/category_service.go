package service

import (
	"my-contacts/internal/model"
	"my-contacts/internal/repository"
)

func FindAllCategories() ([]model.Category, error) {
	return repository.FindAllCategories()
}
