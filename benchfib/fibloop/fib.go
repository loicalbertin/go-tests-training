package fibloop

import "fmt"

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
func PrintFib(maxIt int) {
	for i := 0; i < maxIt; i++ {
		fmt.Printf("%d ", fibonacciLoop(i))
	}
	fmt.Println()
}
