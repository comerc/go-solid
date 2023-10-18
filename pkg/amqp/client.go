package amqp

type RabbitClient struct {
}

func (r RabbitClient) AddMessage() {

}

// ! нарушает DRP, метод следует вынести в отдельную сущность
// func (r RabbitClient) AddToDatabase() { }
