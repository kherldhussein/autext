package autex

// Merges punctuation at the beginning of a word with the previous word
func HandleLead(str []string) []string {
	p := []string{",", ".", "!", "?", ":", ";"}
	for i, s := range str {
		for _, c := range p {
			if string(s[0]) == c && string(s[len(s)-1]) == c && str[len(str)-1] != str[i] {
				str[i-1] = str[i-1] + s
				str = append(str[:i], str[i+1:]...)
			}
		}
	}
	return str
}

// Removes punctuation from the end of a word that isn't the last word
func HandleTrail(str []string) []string {
	p := []string{",", ".", "!", "?", ":", ";"}
	for i, s := range str {
		for _, c := range p {
			if string(s[0]) == c && str[len(str)-1] == str[i] {
				str[i-1] = str[i-1] + s
				str = str[:len(str)-1]
			} else if string(s[0]) == c && string(s[len(s)-1]) != c {
				str[i-1] = str[i-1] + c
				str[i] = s[1:]
			}
		}
	}
	return str
}

// Handles single quotes by merging opening quotes with the following word and removing closing quotes
func HandleQuote(str []string) []string {
	c := 0
	for i, s := range str {
		if string(s[0]) == "'" && c == 0 {
			c++
			str[i+1] = s + str[i+1]
			str = append(str[:i], str[i+1:]...)
		}
	}
	for i, s := range str {
		if s == "'" {
			str[i-1] += s
			str = append(str[:i], str[i+1:]...)
		}
	}
	return str
}
