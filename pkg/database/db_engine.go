package database

// DIP утверждает, что высокоуровневые модули не должны зависеть
// от низкоуровневых модулей, оба должны зависеть от абстракций

// высокоуровневый модуль - реализации интерфейса DbEngine
// низкоуровневый модуль - DbClient
// оба должны зависеть от абстракций - интерфейс DbEngine

// ---

// DIP также утверждает, что абстракции не должны зависеть от деталей,
// но детали должны зависеть от абстракций

// интерфейс и его реализации

type DbEngine interface {
	Add(query string)
	Select(query string)
}

// Реализация интерфейса DbEngine для MySQL

type MySQLEngine struct {
}

func (m MySQLEngine) Add(query string) {

}

func (m MySQLEngine) Select(query string) {

}

// Реализация интерфейса DbEngine для Redis

type RedisEngine struct {
}

func (r RedisEngine) Add(query string) {

}

func (r RedisEngine) Select(query string) {

}
