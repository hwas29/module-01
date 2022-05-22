package beverage

import "tanht/Toppings"

type Beverage struct {
	price 	int
	ice		int
	sugar	int
	tea		int
	milk	string
	size	string
	name	string
	caffein	bool
	recipe 	Preparation
	pricing	toppings.Composite
}

func (this *Beverage) SetRecipe(p Preparation) {
	this.recipe = p
}

func (this *Beverage) GetPrice() int {
	return this.price + this.pricing.GetPrice()
}

func (this *Beverage) SetPrice(p toppings.Composite) {
	this.pricing = p
}

func (this *Beverage) SetPrepare() {
	this.recipe.preparation(this)
}

func (this *Beverage) GetName() string {
	return this.name
}

func (this *Beverage) GetToppings() []string {
	var result []string
	for _, each := range this.pricing.Components {
		result = append(result, each.GetName())
	}
	return result
}

func (this *Beverage) GetNumbersTopping() int {
	return len(this.pricing.Components)
}
