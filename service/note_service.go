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

func (s *NoteService) CreateNote(categoryID int, title string, content string) (model.Note, error) {
	if title == "" {
		return model.Note{}, errors.New("title tidak boleh kosong")
	}
	return s.Repo.Create(categoryID, title, content)
}

func (s *NoteService) GetAllNote() ([]model.Note, error) {
	return s.Repo.GetAll()
}

func (s *NoteService) GetNoteByID(id int) (model.Note, error) {
	return s.Repo.GetByID(id)
}

func (s *NoteService) UpdateNote(id int, categoryID int, title string, content string) (model.Note, error) {
	if title == "" {
		return model.Note{}, errors.New("title tidak boleh kosong")
	}
	return s.Repo.Update(id, categoryID, title, content)
}

func (s *NoteService) DeleteNote(id int) error {
	return s.Repo.Delete(id)
}
