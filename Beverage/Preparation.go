package beverage

type Preparation interface {
	preparation(b *Beverage)
}