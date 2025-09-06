package repository

import (
	"database/sql"

	"github.com/Samith-Ailapperuma/Inventory_Mgt_System/internal/model"
)

type VendorRepository struct {
	DB *sql.DB
}

func NewVendorRepository(db *sql.DB) *VendorRepository {
	return &VendorRepository{DB: db}
}

func (r *VendorRepository) GetAllVendors() ([]model.Vendor, error) {
	rows, err := r.DB.Query("SELECT vendorId, vendorName FROM Vendor")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vendors []model.Vendor
	for rows.Next() {
		var vendor model.Vendor
		if err := rows.Scan(&vendor.Vendor_Id, &vendor.Vendor_Name); err != nil {
			return nil, err
		}
		vendors = append(vendors, vendor)
	}

	return vendors, nil
}

func (r *VendorRepository) CreateVendor(vendor model.Vendor) error {
	_, err := r.DB.Exec("INSERT INTO vendor (vendorId, vendorName) VALUES (?, ?)", vendor.Vendor_Id, vendor.Vendor_Name)
	return err
}
