package api

import (
	"MarcoZillgen/homeChef/internal/storage"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	repo *storage.Repository
}

func NewHandler(repo *storage.Repository) *Handler {
	return &Handler{repo: repo}
}

// helper functions
func (h *Handler) GetParam(r *http.Request, key string) (string, error) {
	vars := mux.Vars(r)
	value, ok := vars[key]
	if !ok {
		return "", fmt.Errorf("missing parameter %s", key)
	}
	return value, nil
}

func (h *Handler) GetParams(r *http.Request, keys ...string) ([]string, error) {
	vars := mux.Vars(r)
	var values []string
	for _, key := range keys {
		value, ok := vars[key]
		if !ok {
			return nil, fmt.Errorf("missing parameter %s", key)
		}
		values = append(values, value)
	}
	return values, nil
}

// handler functions
func (h *Handler) GetItems(w http.ResponseWriter, r *http.Request) {
	items, err := h.repo.GetItems()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(items)
}

func (h *Handler) CreateItem(w http.ResponseWriter, r *http.Request) {
	var item storage.StorageItem
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.repo.CreateItem(item); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) GetItemByID(w http.ResponseWriter, r *http.Request) {
	id, err := h.GetParam(r, "id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	item, err := h.repo.GetItemByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(item)
}

func (h *Handler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	var item storage.StorageItem
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.repo.UpdateItem(item); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	id, err := h.GetParam(r, "id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.repo.DeleteItem(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
