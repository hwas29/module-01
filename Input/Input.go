package input

type Input interface {
	IReader() string
	ISetIo()
}