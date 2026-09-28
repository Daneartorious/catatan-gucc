package handler

import (
	"catatan-backend/model"
	"catatan-backend/service"
	"encoding/json"
	"net/http"
)

type NoteHandler struct {
	Service *service.NoteService
}

func NewNoteHandler(s *service.NoteService) *NoteHandler {
	return &NoteHandler{Service: s}
}

func (h *NoteHandler) Create(w http.ResponseWriter, r *http.Request) {

	var input model.Note
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Data tidak valid", http.StatusBadRequest)
		return
	}

	Note, err := h.Service.CreateNote(input.CategoryID, input.Title, input.Content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Note)
}
