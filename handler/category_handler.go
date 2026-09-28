package handler

import (
	"catatan-backend/model"
	"catatan-backend/service"
	"encoding/json"
	"net/http"
)

type CategoryHandler struct {
	Service *service.CategoryService
}

func NewCategoryHandler(s *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{Service: s}
}

func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {

	var input model.Category
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Data tidak valid", http.StatusBadRequest)
		return
	}

	Category, err := h.Service.CreateCategory(input.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Category)
}

func (h *CategoryHandler) GetAll(w http.ResponseWriter, r *http.Request) {

	category, err := h.Service.GetAllCategory()
	if err != nil {
		http.Error(w, "Gagal mengambil data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(category)
}
