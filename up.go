package autex

import (
	"strconv"
	"strings"
)

// Converts the word before "(up)" to uppercase and removes "(up)".
func Up(s []string) []string {
	for i := 0; i < len(s); i++ {
		if i < 0 {
			continue
		}
		switch {
		case strings.Contains(s[i], "up"):
			if strings.Contains(s[i], "up,") {
				pi, _ := strconv.Atoi(s[i+1][:len(s[i+1])-1])
				for k := i - pi; k < i; k++ {
					s[k] = strings.ToUpper(s[k])
				}
				s = append(s[:i], s[i+2:]...)
			} else {
				s[i-1] = strings.ToUpper(s[i-1])
				s = append(s[:i], s[i+1:]...)
			}
		}
	}
	return s
}
