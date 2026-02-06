// Package hamming
package hamming

import "fmt"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("sequences are not of equal length: %v, %v", len(a), len(b))
	}

	count := 0

	for i := range a {
		if a[i] != b[i] {
			count = count + 1
		}
	}
	return count, nil
}
