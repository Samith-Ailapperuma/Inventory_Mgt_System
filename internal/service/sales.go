package service

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Samith-Ailapperuma/Inventory_Mgt_System/internal/model"
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

func (h *SalesHandler) GetAllItemsInSale(w http.ResponseWriter, r *http.Request) {
	var saleId string
	sales, err := h.Repo.GetAllItemsInSale(saleId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(sales)
}

func (h *SalesHandler) AddItemToSale(w http.ResponseWriter, r *http.Request) {
	var saleItem model.SaleItem
	if err := json.NewDecoder(r.Body).Decode(&saleItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var saleId string

	if saleItem.Sale_Id == "" {
		sale := model.Sale{
			Sales_Date:  time.Now(),
			Sale_Amount: 0,
		}
		saleId, err := h.Repo.CreateSale(sale)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		saleItem.Sale_Id = saleId
	} else {
		saleId = saleItem.Sale_Id
	}

	err := h.Repo.AddItemToSale(saleItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.Repo.UpdateSaleAmount(saleItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"saleId": saleId})
}
