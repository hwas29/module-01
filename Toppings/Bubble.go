package toppings


type Bubble struct {
	Topping
}

func NewBubble() Component {
	return &Bubble{
		Topping: Topping{
			name: "Bubble",
			price: 10,
		},
	}
	
}

