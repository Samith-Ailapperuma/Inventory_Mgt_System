package model

import "time"

type Sale struct {
	Sale_Id     string
	Sales_Date  time.Time
	Sale_Amount float32
}
