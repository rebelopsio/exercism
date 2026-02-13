// Package isogram
package isogram

import "strings"

func IsIsogram(word string) bool {
	seen := make(map[rune]bool)
	for _, v := range strings.ToLower(word) {
		if v == ' ' || v == '-' {
			continue
		}
		if _, ok := seen[v]; ok {
			return false
		} else {
			seen[v] = true
		}
	}

	return true
}
