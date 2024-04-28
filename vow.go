package autex

// Replaces "a" with "an" if the next word starts with a vowel or "h"
func Vow(str []string) []string {
	vw := []string{"a", "e", "i", "o", "u", "h", "A", "E", "I", "O", "U", "H"}
	for i := 0; i < len(str); i++ {
		s := str[i]
		for _, v := range vw {
			if s == "a" && len(str[i+1]) > 0 && string(str[i+1][0]) == v {
				str[i] = "an"
			} else if s == "A" && len(str[i+1]) > 0 && string(str[i+1][0]) == v {
				str[i] = "An"
			}
		}
	}
	return str
}
