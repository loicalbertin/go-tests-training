package fibloop

import (
	"bytes"
	"fmt"
)

func fibonacciLoop(n int) int {
	if n < 2 {
		return n
	}

	a, b := 1, 0

	for n > 2 {
		n, a, b = n-1, a+b, a
	}

	return a + b
}

// PrintFib prints a fibonacci sequence up to maxIt iterations
func PrintFib(maxIt int) string {
	b := bytes.Buffer{}
	for i := 0; i < maxIt; i++ {
		b.WriteString(fmt.Sprintf("%d ", fibonacciLoop(i)))
	}
	return b.String()
}
