package basic

import (
	"math"
)

// Max is a dummy function that allows to get the max of 2 integers
func Max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
