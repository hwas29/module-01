package toppings

type Component interface {
	GetPrice() int
	GetName() string
}