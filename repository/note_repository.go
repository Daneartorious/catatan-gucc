package repository

import (
	"catatan-backend/model"
	"database/sql"
	"errors"
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

func (r *NoteRepository) GetAll() ([]model.Note, error) {
	query := `
	SELECT n.id, n.category_id, n.title, n.content, c.name AS category_name
	FROM notes n 
	LEFT JOIN categories c ON n.category_id = c.id
`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	note := make([]model.Note, 0)
	for rows.Next() {
		var n model.Note

		if err := rows.Scan(&n.ID, &n.CategoryID, &n.Title, &n.Content, &n.CategoryName); err != nil {
			return nil, err
		}
		note = append(note, n)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return note, nil
}

func (r *NoteRepository) GetByID(id int) (model.Note, error) {
	var n model.Note
	query := `
	SELECT n.id, n.category_id, n.title, n.content, c.name AS category_name
	FROM notes n 
	LEFT JOIN categories c ON n.category_id = c.id
	WHERE n.id = $1
	`
	err := r.DB.QueryRow(query, id).Scan(&n.ID, &n.CategoryID, &n.Title, &n.Content, &n.CategoryName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return n, errors.New("note tidak ditemukan")
		}
		return n, err
	}
	return n, nil
}

func (r *NoteRepository) Update(id int, categoryID int, title string, content string) (model.Note, error) {
	query := `UPDATE notes SET category_id = $1, title = $2, content = $3 WHERE id = $4`

	var updatedNote model.Note
	result, err := r.DB.Exec(query, categoryID, title, content, id)
	if err != nil {
		return updatedNote, err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return updatedNote, errors.New("note tidak ditemukan")
	}
	updatedNote.ID = id
	updatedNote.CategoryID = &categoryID
	updatedNote.Title = title
	updatedNote.Content = content

	return updatedNote, nil
}

func (r *NoteRepository) Delete(id int) error {
	query := `DELETE FROM notes WHERE id = $1`
	result, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("note tidak ditemukan")
	}
	return nil
}
