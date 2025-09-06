package service

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Samith-Ailapperuma/Inventory_Mgt_System/internal/repository"
)

type Sales struct {
	Sale_Id     string
	Sales_Date  time.Time
	Sale_Amount float32
}

type SalesHandler struct {
	Repo *repository.SalesRepository
}

func NewSalesHandler(repo *repository.SalesRepository) *SalesHandler {
	return &SalesHandler{Repo: repo}
}

func (h *SalesHandler) GetAllSales(w http.ResponseWriter, r *http.Request) {
	sales, err := h.Repo.GetAllSales()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(sales)
}
