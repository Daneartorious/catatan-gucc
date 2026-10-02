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
		http.Error(w, "data tidak valid", http.StatusBadRequest)
		return
	}

	if input.CategoryID == nil {
		http.Error(w, "category_id tidak boleh kosong", http.StatusBadRequest)
		return
	}
	note, err := h.Service.CreateNote(*input.CategoryID, input.Title, input.Content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(note)
}

func (h *NoteHandler) GetAll(w http.ResponseWriter, r *http.Request) {

	note, err := h.Service.GetAllNote()
	if err != nil {
		http.Error(w, "gagal mengambil data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(note)
}

func (h *NoteHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "id harus berupa angka positif", http.StatusBadRequest)
		return
	}

	note, err := h.Service.GetNoteByID(id)
	if err != nil {
		if err.Error() == "note tidak ditemukan" {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(note)
}

func (h *NoteHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "id harus berupa angka positif", http.StatusBadRequest)
		return
	}
	var updatedNote model.Note
	err = json.NewDecoder(r.Body).Decode(&updatedNote)
	if err != nil {
		http.Error(w, "data tidak valid", http.StatusBadRequest)
		return
	}

	if updatedNote.CategoryID == nil {
		http.Error(w, "category_id tidak boleh kosong", http.StatusBadRequest)
		return

	}
	if updatedNote.Content == "" {
		http.Error(w, "content tidak boleh kosong", http.StatusBadRequest)
		return
	}

	result, err := h.Service.UpdateNote(id, *updatedNote.CategoryID, updatedNote.Title, updatedNote.Content)
	if err != nil {
		if err.Error() == "note tidak ditemukan" {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (h *NoteHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "id harus berupa angka positif", http.StatusBadRequest)
		return
	}
	err = h.Service.DeleteNote(id)
	if err != nil {
		if err.Error() == "note tidak ditemukan" {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Note berhasil dihapus"})
}
