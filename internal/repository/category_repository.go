package repository

import (
	"context"
	"my-contacts/internal/infra"
	"my-contacts/internal/model"
)

func FindAllCategories() ([]model.Category, error) {
	rows, err := infra.DB.Query(
		context.Background(),
		"SELECT id, name FROM categories",
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories = []model.Category{}

	for rows.Next() {
		var c model.Category
		err := rows.Scan(&c.ID, &c.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}

	return categories, nil
}
