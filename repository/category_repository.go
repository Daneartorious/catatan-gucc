package repository

import (
	"catatan-backend/model"
	"database/sql"
)

type CategoryRepository struct {
	DB *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{DB: db}
}

func (r *CategoryRepository) Create(name string) (model.Category, error) {
	var n model.Category
	query := `INSERT INTO categories (name) 
	VALUES ($1) RETURNING id, name`
	err := r.DB.QueryRow(query, name).Scan(
		&n.ID,
		&n.Name,
	)
	return n, err
}

func (r *CategoryRepository) GetAll() ([]model.Category, error) {
	query := `SELECT id, name FROM categories`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var category []model.Category
	for rows.Next() {
		var c model.Category

		if err := rows.Scan(&c.ID, &c.Name); err != nil {
			return nil, err
		}
		category = append(category, c)
	}
	return category, nil
}
