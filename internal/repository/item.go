package repository

import (
	"database/sql"

	"github.com/Samith-Ailapperuma/Inventory_Mgt_System/internal/model"
)

type ItemRepository struct {
	DB *sql.DB
}

func (r *ItemRepository) GetAllItems() ([]model.Item, error) {
	rows, err := r.DB.Query("SELECT * FROM Items")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []model.Item
	for rows.Next() {
		var item model.Item
		if err := rows.Scan(&item.Item_Id, &item.Item_Name, &item.Unit_Price, &item.Vendor_Id); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}
