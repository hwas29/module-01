package media


type Media struct {
	components []Notify
}

func (this *Media) SendMessage(message string) {
	for _, each := range this.components {
		each.SendMessage(message)
	}
}

func (this *Media) AddMedia(c Notify) {
	this.components = append(this.components, c)
}