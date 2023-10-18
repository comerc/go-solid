package models

type Popato struct {
	Price float64
}

func (p Popato) GetPrice() float64 {
	return 100
}
