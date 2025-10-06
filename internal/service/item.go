package service

import (
	"encoding/json"
	"net/http"

	"github.com/Samith-Ailapperuma/Inventory_Mgt_System/internal/model"
	"github.com/Samith-Ailapperuma/Inventory_Mgt_System/internal/repository"
)

type Item struct {
	Item_Id    string
	Vendor_Id  string
	Item_Name  string
	Unit_Price float32
}

type ItemHandler struct {
	Repo *repository.ItemRepository
}

func NewItemHandler(repo *repository.ItemRepository) *ItemHandler {
	return &ItemHandler{Repo: repo}
}

func (h *ItemHandler) AddNewItem(w http.ResponseWriter, r *http.Request) {
	var item model.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.Repo.AddNewItem(item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *ItemHandler) GetAllItems(w http.ResponseWriter, r *http.Request) {
	items, err := h.Repo.GetAllItems()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(items)
}
