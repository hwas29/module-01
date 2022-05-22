package input

import (
	"bufio"
	"os"
)

type Console struct {
	reader *bufio.Reader
}

func (this *Console) ISetIo() {

	this.reader = bufio.NewReader(os.Stdin)

}

func (this *Console) IReader() string {

	result, _ := this.reader.ReadString('\n')
	result = result[:len(result)-1]

	return result
}
