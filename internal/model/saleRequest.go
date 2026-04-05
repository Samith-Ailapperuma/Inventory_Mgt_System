package model

import "time"

type SaleRequest struct {
	Sales_Date time.Time
	Items      []SaleItemInput
}
