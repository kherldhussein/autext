package autex

import (
	"strconv"
	"strings"
)

// Capitalizes the word before "(cap)" and removes "(cap)".
func WordZ(s []string) []string {
	for i := 0; i < len(s); i++ {
		if i <= 0 {
			continue
		}
		switch {
		case strings.Contains(s[i], "cap"):
			if strings.Contains(s[i], "cap,") {
				pi, _ := strconv.Atoi(s[i+1][:len(s[i+1])-1])
				for k := i - pi; k < i; k++ {
					s[k] = strings.Title(s[k])
				}
				s = append(s[:i], s[i+2:]...)
			} else {
				s[i-1] = strings.Title(s[i-1])
				s = append(s[:i], s[i+1:]...)
			}
		}
	}
	return s
}
