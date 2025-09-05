package raindrops

import "fmt"

func Convert(number int) string {
	ans := ""
	if number%3 == 0 {
		ans = ans + "Pling"
	}
	if number%5 == 0 {
		ans = ans + "Plang"
	}
	if number%7 == 0 {
		ans = ans + "Plong"
	}
	if ans == "" {
		ans = fmt.Sprintf("%d", number)
	}
	return ans
}
