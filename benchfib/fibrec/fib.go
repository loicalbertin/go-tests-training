package fibrec

import (
	"bytes"
	"fmt"
)

func fibonacciRecursion(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacciRecursion(n-1) + fibonacciRecursion(n-2)
}

// PrintFib prints a fibonacci sequence up to maxIt iterations
func PrintFib(maxIt int) string {
	b := bytes.Buffer{}
	for i := 0; i < maxIt; i++ {
		b.WriteString(fmt.Sprintf("%d ", fibonacciRecursion(i)))
	}
	return b.String()
}
