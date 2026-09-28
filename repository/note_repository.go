package repository

import (
	"catatan-backend/model"
	"database/sql"
)

type NoteRepository struct {
	DB *sql.DB
}

func NewNoteRepository(db *sql.DB) *NoteRepository {
	return &NoteRepository{DB: db}
}

func (r *NoteRepository) Create(category_id int, title string, content string) (model.Note, error) {
	var n model.Note
	query := `INSERT INTO notes (category_id, title, content) 
	VALUES ($1, $2, $3) RETURNING id, category_id, title, content`
	err := r.DB.QueryRow(query, category_id, title, content).Scan(
		&n.ID,
		&n.CategoryID,
		&n.Title,
		&n.Content,
	)
	return n, err
}
