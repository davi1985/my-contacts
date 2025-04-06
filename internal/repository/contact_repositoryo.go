package repository

import (
	"context"
	"my-contacts/internal/infra"
	"my-contacts/internal/model"
)

func FindAllContacts() ([]model.Contact, error) {
	rows, err := infra.DB.Query(
		context.Background(),
		`SELECT id, name, email, phone, category_id FROM contacts`,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contacts = []model.Contact{}

	for rows.Next() {
		var c model.Contact
		err := rows.Scan(&c.ID, &c.Name, &c.Email, &c.Phone, &c.CategoryID)
		if err != nil {
			return nil, err
		}
		contacts = append(contacts, c)
	}

	return contacts, nil
}
