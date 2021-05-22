package product

import "time"

//Modelo del producto
type Model struct {
	ID uint
	Name string
	Observations string
	Price int
	CreateAt time.Time
	UpdateAt time.Time
}
//slice de model
type Models []* Model

type Storage interface {
	Create(*Model) error
	Update(*Model) error
	GetAll()(Models, error)
	GetByID(uint) (*Model, error)
	Delete(uint) error
}