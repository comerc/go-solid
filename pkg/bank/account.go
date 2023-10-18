package bank

// LSP - это абстракция "суперкласса", т.е. interface на Go

type Account interface {
	Withdraw(money float64)
}

type DefaultAccount struct {
	Balance float64
}

func (acc DefaultAccount) Withdraw(money float64) {

}
