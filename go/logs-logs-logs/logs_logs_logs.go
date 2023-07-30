package logs

import (
	"strings"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	// rune's must be surrounded by single quotes
	m := map[rune]string{'❗': "recommendation", '🔍': "search", '☀': "weather"}
	i := 9223372036854775807
	ans := "default"

	for k, v := range m {
		index := strings.IndexRune(log, k)
		if index > -1 && index < i {
			i = index
			ans = v
		}
	}

	return ans
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	r := []rune(log)

	for i := 0; i < len(r); i++ {
		if r[i] == oldRune {
			r[i] = newRune
		}
	}

	ans := string(r)
	return ans
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
