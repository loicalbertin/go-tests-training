package fibrec

import "fmt"

func fibonacciRecursion(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacciRecursion(n-1) + fibonacciRecursion(n-2)
}

// PrintFib prints a fibonacci sequence up to maxIt iterations
func PrintFib(maxIt int) {
	for i := 0; i < maxIt; i++ {
		fmt.Printf("%d ", fibonacciRecursion(i))
	}
	fmt.Println()
}
