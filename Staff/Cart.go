package staff

import (
	"fmt"
	beverage "tanht/Beverage"
)

type Cart struct {
	cart []*beverage.Beverage
}

func (this *Cart) AddToCart(e *beverage.Beverage) {
	this.cart = append(this.cart, e)
}

func (this *Cart) GetPrice() int {
	result := 0
	for _, each := range this.cart {
		result += each.GetPrice()
	}
	return result
}

func (this *Cart) ShowBeverage() {
	for index, each := range this.cart {
		if each.GetNumbersTopping() == 0 {
			fmt.Print(index+1, " : ", each.GetName(), " with no toppings")
		} else {
			fmt.Print(index+1, " : ", each.GetName(), " with toppings ")
		}
		toppings := each.GetToppings()
		for i, name := range toppings {
			fmt.Print(name)
			if i < len(toppings)-1 {
				fmt.Print(", ")
			}
		}
		fmt.Println(".")
	}
}

func (this *Cart) GetNameBeverages() []string {

	var result []string
	for _, each := range this.cart {
		result = append(result, each.GetName())
	}
	return result
}

func (this *Cart) GetNumbers() int {
	return len(this.cart)
}
