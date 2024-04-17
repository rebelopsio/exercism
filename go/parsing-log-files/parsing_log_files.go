package parsinglogfiles

import (
	"fmt"
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	m, _ := regexp.MatchString(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\].*`, text)
	return m
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[-*~=]*>`)
	sl := re.Split(text, -1)
	return sl
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)"(.*password.*)"`)
	sl := make([]string, 0)
	for _, line := range lines {
		if re.FindString(line) != "" {
			sl = append(sl, re.FindString(line))
		}
	}
	return len(sl)
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`(end-of-line\d*)`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`(User\s+\w*)`)
	for i, line := range lines {
		if re.FindString(line) != "" {
			ru := regexp.MustCompile(` (\w)+`)
			username := strings.TrimSpace(ru.FindString(re.FindString(line)))
			lines[i] = fmt.Sprintf("[USR] %s %s", username, line)
		}
	}
	return lines
}
