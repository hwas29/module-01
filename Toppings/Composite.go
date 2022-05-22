package toppings

import "fmt"

type Composite struct {
	Components []Component
}

func (this *Composite) GetPrice() int {
	result := 0
	for _, each := range this.Components {
		result += each.GetPrice()
	}
	return result
}

func (this *Composite) AddComponent(c Component) {
	this.Components = append(this.Components, c)
}

func (this *Composite) ShowComponent() {
	for _, each := range this.Components {
		fmt.Print(each.GetName(), " ")
	}
	fmt.Println()
}
