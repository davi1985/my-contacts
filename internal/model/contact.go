package model

import "github.com/google/uuid"

type Contact struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	CategoryID uuid.UUID `json:"category_id"`
}

func NewContact(name, email, phone string, categoryID uuid.UUID) Contact {
	return Contact{
		ID:         uuid.New(),
		Name:       name,
		Email:      email,
		Phone:      phone,
		CategoryID: categoryID,
	}
}
