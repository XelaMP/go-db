package invoiceheader

import "time"
//Model de invoice header
type Model struct {
	ID uint
	Client string
	CreateAt time.Time
	UpdateAt time.Time
}
