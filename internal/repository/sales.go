package repository

import (
	"database/sql"
	"fmt"

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

func (r *SalesRepository) GetAllItemsInSale(saleId string) ([]model.SaleItemDetails, error) {
	rows, err := r.DB.Query("SELECT s.Quantity_Sold, i.item_name, i.unit_price FROM SalesItem s JOIN Item i ON i.item_id=s.item_id where s.sale_id=?", saleId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var saleItems []model.SaleItemDetails

	for rows.Next() {
		var saleItem model.SaleItemDetails
		if err := rows.Scan(&saleItem.Qty_Sold, &saleItem.Item_Name, &saleItem.Unit_Price); err != nil {
			return nil, err
		}
		saleItem.Total_Price = saleItem.Unit_Price * float32(saleItem.Qty_Sold)
		saleItems = append(saleItems, saleItem)
	}

	return saleItems, nil
}

func (r *SalesRepository) CreateSale(sale model.Sale) (string, error) {
	result, err := r.DB.Exec("INSERT INTO sales (Sales_Date, Sales_Amount) VALUES (?, ?)", sale.Sales_Date, sale.Sale_Amount)
	if err != nil {
		return "", err
	}

	id, _ := result.LastInsertId()
	return fmt.Sprintf("%d", id), nil
}

func (r *SalesRepository) AddItemToSale(saleItem model.SaleItem) error {
	_, err := r.DB.Exec("INSERT INTO salesItem (Sales_Id, Item_Id, Quantity_Sold) VALUES (?, ?, ?)", saleItem.Sale_Id, saleItem.Item_Id, saleItem.Qty_Sold)
	return err
}

func (r *SalesRepository) UpdateSaleAmount(saleItem model.SaleItem) error {
	var UnitPrice float32
	err := r.DB.QueryRow("SELECT Unit_Price FROM items WHERE Item_Id = ?", saleItem.Item_Id).Scan(&UnitPrice)
	if err != nil {
		return err
	}

	itemTotal := float32(saleItem.Qty_Sold) * UnitPrice

	_, err = r.DB.Exec(
		"UPDATE sales SET Sales_Amount = Sales_Amount + ? WHERE Sales_Id = ?",
		itemTotal, saleItem.Sale_Id,
	)
	return err
}
