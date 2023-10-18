package database

type DbClient struct {
}

func (db DbClient) AddToDatabase(engine DbEngine) {

}

type Item struct {
	Name string
}

// для OCP, выносим реализацию row
type Product interface {
	GetPrice() float64
}

func (db DbClient) ReadTable() float64 {
	var row = Item{
		Name: "Картошка",
	}
	var price = .0
	// ! нарушает OCP, новые типы рядов следует добавлять снаружи
	// ! не изменяя код этого метода, т.е. Product(row).GetPrice()
	switch row.Name {
	case "Картошка":
		price = 100
	case "Чипсы":
		price = 10
	}
	return price
}
