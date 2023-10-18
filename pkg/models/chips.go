package models

type Chips struct {
	Price float64
}

func (p Chips) GetPrice() float64 {
	return 10
}
