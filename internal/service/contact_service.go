package service

import (
	"my-contacts/internal/model"
	"my-contacts/internal/repository"
)

func FindAllContacts() ([]model.Contact, error) {
	return repository.FindAllContacts()
}
