package beverage


type BubbleTea struct {
}


func (this *BubbleTea) preparation(b *Beverage) {

	b.name = "BubbleTea"
	b.ice = 100
	b.milk = "Sua Long Thanh"
	b.caffein = false
	b.price = 30
	b.sugar = 70
	b.tea = 50
	b.size = "M"

}
