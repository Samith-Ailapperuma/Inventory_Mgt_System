package repository

import (
	"database/sql"
	"fmt"

	"github.com/Samith-Ailapperuma/Inventory_Mgt_System/internal/model"
)

type ItemRepository struct {
	DB *sql.DB
}

func NewItemRepository(db *sql.DB) *ItemRepository {
	return &ItemRepository{DB: db}
}

func (r *ItemRepository) AddNewItem(item model.Item) (string, error) {
	result, err := r.DB.Exec("INSERT INTO item (Vendor_Id, Item_Name, Unit_Price) VALUES (?, ?, ?)", item.Vendor_Id, item.Item_Name, item.Unit_Price)
	if err != nil {
		return "", err
	}

	id, _ := result.LastInsertId()
	return fmt.Sprintf("%d", id), nil
}

func (r *ItemRepository) GetAllItems() ([]model.Item, error) {
	rows, err := r.DB.Query("SELECT * FROM Item")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []model.Item
	for rows.Next() {
		var item model.Item
		if err := rows.Scan(&item.Item_Id, &item.Vendor_Id, &item.Item_Name, &item.Unit_Price); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}
