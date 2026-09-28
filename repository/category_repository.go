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
	query := `INSERT INTO Categories (name) 
	VALUES ($1) RETURNING id, name`
	err := r.DB.QueryRow(query, name).Scan(
		&n.ID,
		&n.Name,
	)
	return n, err
}
