package autex

import (
	"fmt"
	"strconv"
)

// Converts hexadecimal and binary numbers to decimal
func Hexbin(str []string) []string {
	for i := 0; i < len(str); i++ {
		switch str[i] {
		case "(hex)":
			if num, err := strconv.ParseInt(str[i-1], 16, 64); err == nil {
				str[i-1] = fmt.Sprint(num)
				str = append(str[:i], str[i+1:]...)
				i--
			}
		case "(bin)":
			if num, err := strconv.ParseInt(str[i-1], 2, 64); err == nil {
				str[i-1] = fmt.Sprint(num)
				str = append(str[:i], str[i+1:]...)
				i--
			}
		}
	}
	return str
}
