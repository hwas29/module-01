package media

import (
	"fmt"
	"sync"
)

var onceMessenger sync.Once

type Messenger struct {
}

var InstanceMessenger *Messenger

func GetInstanceMessenger() (*Messenger, int) {
	if InstanceMessenger == nil {
		onceMessenger.Do(
			func() {
				InstanceMessenger = &Messenger{}
			},
		)
		return InstanceMessenger, 0
	}

	fmt.Println("Messenger has been added already")
	return InstanceMessenger, -1
}

func (this *Messenger) SendMessage(message string) {
	fmt.Println("Messenger: ", message)
}
