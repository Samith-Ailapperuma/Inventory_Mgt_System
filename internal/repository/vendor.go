package repository

import (
	"database/sql"
	"fmt"

	"github.com/Samith-Ailapperuma/Inventory_Mgt_System/internal/model"
)

type VendorRepository struct {
	DB *sql.DB
}

func NewVendorRepository(db *sql.DB) *VendorRepository {
	return &VendorRepository{DB: db}
}

func (r *VendorRepository) GetAllVendors() ([]model.Vendor, error) {
	rows, err := r.DB.Query("SELECT Vendor_Id, Vendor_Name, Vendor_Address FROM Vendor")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vendors []model.Vendor
	for rows.Next() {
		var vendor model.Vendor
		if err := rows.Scan(&vendor.Vendor_Id, &vendor.Vendor_Name, &vendor.Vendor_Address); err != nil {
			return nil, err
		}
		vendors = append(vendors, vendor)
	}

	return vendors, nil
}

func (r *VendorRepository) CreateVendor(vendor model.Vendor) (string, error) {
	result, err := r.DB.Exec("INSERT INTO Vendor (Vendor_Name, Vendor_Address) VALUES (?, ?)", vendor.Vendor_Name, vendor.Vendor_Address)
	if err != nil {
		return "", err
	}

	id, _ := result.LastInsertId()
	return fmt.Sprintf("%d", id), nil
}
