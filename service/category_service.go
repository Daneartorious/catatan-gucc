package service

import (
	"catatan-backend/model"
	"catatan-backend/repository"
	"errors"
)

type CategoryService struct {
	Repo *repository.CategoryRepository
}

func NewCategoryService(repo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{Repo: repo}
}

func (s *CategoryService) CreateCategory(name string) (model.Category, error) {
	if name == "" {
		return model.Category{}, errors.New("nama tidak boleh kosong")
	}
	return s.Repo.Create(name)
}

func (s *CategoryService) GetAllCategory() ([]model.Category, error) {
	return s.Repo.GetAll()
}

func (s *CategoryService) DeleteCategory(id int) error {
	return s.Repo.Delete(id)
}