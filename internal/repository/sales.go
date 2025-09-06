package repository

import (
	"database/sql"

	"github.com/Samith-Ailapperuma/Inventory_Mgt_System/internal/model"
)

type SalesRepository struct {
	DB *sql.DB
}

func NewSalesRepository(db *sql.DB) *SalesRepository {
	return &SalesRepository{DB: db}
}

func (r *SalesRepository) GetAllSales() ([]model.Sale, error) {
	rows, err := r.DB.Query("SELECT * FROM Sales")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sales []model.Sale
	for rows.Next() {
		var sale model.Sale
		if err := rows.Scan(&sale.Sale_Id, &sale.Sale_Amount, &sale.Sales_Date); err != nil {
			return nil, err
		}
		sales = append(sales, sale)
	}

	return sales, nil
}
