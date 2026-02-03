// Package collatzconjecture
package collatzconjecture

import "fmt"

func CollatzConjecture(n int) (int, error) {
	if n < 1 {
		return 0, fmt.Errorf("%d is invalid", n)
	}
	ans := 0
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
			ans = ans + 1
		} else {
			n = n*3 + 1
			ans = ans + 1
		}
	}
	return ans, nil
}
