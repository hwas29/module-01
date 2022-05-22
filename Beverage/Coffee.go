package beverage

type Coffee struct {
}
func (this *Coffee) preparation(b *Beverage) {

	b.name = "Coffee"
	b.ice = 100
	b.milk = "Sua Ong Tho"
	b.caffein = true
	b.price = 20
	b.sugar = 50
	b.tea = 0
	b.size = "M"

}
