package fibchan

import (
	"bytes"
	"fmt"
)

func fibChan(maxIt, buffer uint) chan int {
	ch := make(chan int, buffer)
	a, b := 0, 1
	go func() {
		for i := uint(0); i < maxIt; i++ {
			ch <- a
			a, b = b, a+b
		}
		close(ch)
	}()
	return ch
}

// PrintFib prints a fibonacci sequence up to maxIt iterations
// Buffer allows to control the chanel buffering
func PrintFib(maxIt, buffer uint) string {
	b := bytes.Buffer{}
	for i := range fibChan(maxIt, buffer) {
		b.WriteString(fmt.Sprintf("%d ", i))
	}
	return b.String()
}
