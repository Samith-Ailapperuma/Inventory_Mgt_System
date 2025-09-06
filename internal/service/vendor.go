package service

import (
	"encoding/json"
	"net/http"

	"github.com/Samith-Ailapperuma/Inventory_Mgt_System/internal/model"
	"github.com/Samith-Ailapperuma/Inventory_Mgt_System/internal/repository"
)

type VendorHandler struct {
	Repo *repository.VendorRepository
}

func NewVendorHandler(repo *repository.VendorRepository) *VendorHandler {
	return &VendorHandler{Repo: repo}
}

func (h *VendorHandler) GetVendors(w http.ResponseWriter, r *http.Request) {
	vendors, err := h.Repo.GetAllVendors()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(vendors)
}

func (h *VendorHandler) CreateVendor(w http.ResponseWriter, r *http.Request) {
	var vendor model.Vendor
	if err := json.NewDecoder(r.Body).Decode(&vendor); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.Repo.CreateVendor(vendor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
