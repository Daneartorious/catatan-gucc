package handler

import (
	"catatan-backend/model"
	"catatan-backend/service"
	"encoding/json"
	"net/http"
	"strconv"
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

	note, err := h.Service.CreateNote(*input.CategoryID, input.Title, input.Content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(note)
}

func (h *NoteHandler) GetAll(w http.ResponseWriter, r *http.Request) {

	note, err := h.Service.GetAllNote()
	if err != nil {
		http.Error(w, "Gagal mengambil data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(note)
}

func (h *NoteHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
    	http.Error(w, "ID tidak valid", http.StatusBadRequest)
    	return
	}
	err = h.Service.DeleteNote(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Note berhasil dihapus"})
}
