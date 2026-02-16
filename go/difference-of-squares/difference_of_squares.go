// Package diffsquares
package diffsquares

import "math"

func SquareOfSum(n int) int {
	sum := 0
	for i := range n + 1 {
		sum += i
	}
	return int(math.Pow(float64(sum), 2.0))
}

func SumOfSquares(n int) int {
	sum := 0
	for i := range n + 1 {
		i = int(math.Pow(float64(i), 2.0))
		sum += i
	}
	return sum
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
