package repository

import (
	"real-time-forum/backend/models"
	"real-time-forum/database"
)

func GetCategories() ([]models.CategoryResponse, error) {
	rows, err := database.DB.Query(`
		SELECT
			id,
			name
		FROM categories
		ORDER BY name
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.CategoryResponse

	for rows.Next() {

		var category models.CategoryResponse

		err := rows.Scan(
			&category.ID,
			&category.Name,
		)
		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if categories == nil {
		categories = []models.CategoryResponse{}
	}

	return categories, nil
}
