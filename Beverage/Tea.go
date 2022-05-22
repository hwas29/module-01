package beverage



type Tea struct {
}



func (this *Tea) preparation(b *Beverage) {

	b.name = "Tea"
	b.ice = 100
	b.milk = "None"
	b.caffein = true
	b.price = 15
	b.sugar = 0
	b.tea = 80
	b.size = "M"

}
