package model

type SaleItemDetails struct {
	Sale_Id     string
	Qty_Sold    int
	Item_Id     string
	Vendor_Id   string
	Item_Name   string
	Unit_Price  float32
	Total_Price float32
}
