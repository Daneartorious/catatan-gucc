package repository

import (
	"catatan-backend/model"
	"database/sql"
	"errors"
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

func (r *CategoryRepository) Update(id int, name string) (model.Category, error) {
	query := `UPDATE categories SET name = $1 WHERE id = $2`
	var updatedCategory model.Category
	result, err := r.DB.Exec(query, name, id)
	if err != nil {
		return updatedCategory, err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return updatedCategory, errors.New("category tidak ditemukan")
	}
	updatedCategory.ID = id
	updatedCategory.Name = name

	return updatedCategory, nil
}

func (r *CategoryRepository) Delete(id int) error {
	query := `DELETE FROM categories WHERE id = $1`
	result, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("category tidak ditemukan")
	}
	return nil
}
