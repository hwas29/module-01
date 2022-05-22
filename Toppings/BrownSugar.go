package toppings

type BrownSugar struct {
	Topping
}

func NewBrownSugar() Component {
	return &BrownSugar{
		Topping: Topping{
			name:  "BrownSugar",
			price: 20,
		},
	}
}
