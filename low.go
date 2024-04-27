package autex

import (
	"strconv"
	"strings"
)

// Converts the word before "(low)" to lowercase and removes "(low)".
func Low(s []string) []string {
	for i := 0; i < len(s); i++ {
		if i <= 0 {
			continue
		}
		switch {
		case strings.Contains(strings.ToLower(s[i]), "low"):
			if strings.Contains(s[i], "low,") {
				pi, _ := strconv.Atoi(s[i+1][:len(s[i+1])-1])
				for k := i - pi; k < i; k++ {
					s[k] = strings.ToLower(s[k])
				}
				s = append(s[:i], s[i+2:]...)
			} else {
				s[i-1] = strings.ToLower(s[i-1])
				s = append(s[:i], s[i+1:]...)
			}
		}
	}
	return s
}
