package model

import "github.com/google/uuid"

type Category struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func NewCategory(name string) Category {
	return Category{
		ID:   uuid.New(),
		Name: name,
	}
}
