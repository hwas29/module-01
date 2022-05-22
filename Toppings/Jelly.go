package toppings


type Jelly struct {
	Topping
}

func NewJelly() Component {
	return &Jelly{
		Topping: Topping{
			name: "Jelly",
			price: 10,
		},
	}
}


