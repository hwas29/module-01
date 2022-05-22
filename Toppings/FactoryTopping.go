package toppings

type Topping struct {
	name  string
	price int
}

func (this *Topping) GetName() string {
	return this.name
}

func (this *Topping) GetPrice() int {
	return this.price
}

func FactoryTopping(typeTopping string) (Component, int) {
	if typeTopping == "Bubble" {
		return NewBubble(), 0
	}
	if typeTopping == "Jelly" {
		return NewJelly(), 0
	}
	if typeTopping == "BrownSugar" {
		return NewBrownSugar(), 0
	}
	return nil, -1
}
