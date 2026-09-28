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

func (r *NoteRepository) Create(Note_id int, title string, content string) (model.Note, error) {
	var n model.Note
	query := `INSERT INTO notes (Note_id, title, content) 
	VALUES ($1, $2, $3) RETURNING id, Note_id, title, content`
	err := r.DB.QueryRow(query, Note_id, title, content).Scan(
		&n.ID,
		&n.CategoryID,
		&n.Title,
		&n.Content,
	)
	return n, err
}

func (r *NoteRepository) GetAll() ([]model.Note, error) {
	query := `SELECT id, category_id, title, content FROM notes`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var note []model.Note
	for rows.Next() {
		var n model.Note

		if err := rows.Scan(&n.ID, &n.CategoryID, &n.Title, &n.Content); err != nil {
			return nil, err
		}
		note = append(note, n)
	}
	return note, nil
}
