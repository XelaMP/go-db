package invoiceitem

import "time"
//Model de invoice item
type Model struct {
	ID uint
	InvoiceHeaderID uint
	ProductID uint
	CreatedAt time.Time
	UpdateAt time.Time
}

