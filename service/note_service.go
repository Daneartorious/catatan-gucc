package service

import (
	"catatan-backend/model"
	"catatan-backend/repository"
	"errors"
)

type NoteService struct {
	Repo *repository.NoteRepository
}

func NewNoteService(repo *repository.NoteRepository) *NoteService {
	return &NoteService{Repo: repo}
}

func (s *NoteService) CreateNote(category_id int, title string, content string) (model.Note, error) {
	if title == "" {
		return model.Note{}, errors.New("title tidak boleh kosong")
	}
	return s.Repo.Create(category_id, title, content)
}

func (s *NoteService) GetAllNote() ([]model.Note, error) {
	return s.Repo.GetAll()
}
