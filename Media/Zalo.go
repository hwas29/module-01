package media

import (
	"fmt"
	"sync"
)

var onceZalo sync.Once

type Zalo struct {
}

var InstanceZalo *Zalo

func GetInstanceZalo() (*Zalo, int) {
	if InstanceZalo == nil {
		onceZalo.Do(
			func() {
				InstanceZalo = &Zalo{}
			},
		)
		return InstanceZalo, 0
	}

	return InstanceZalo, -1
}

func (this *Zalo) SendMessage(message string) {
	fmt.Println("Zalo: ", message)
}
