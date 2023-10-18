package os

// ISP - один метод на интерфейс (в идеале)

type Opener interface {
	Open()
}
