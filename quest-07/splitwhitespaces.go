package piscine

func SplitWhiteSpaces(s string) []string {
	word := ""
	var strings []string
	for i := 0; i < len(s); i++ {
		c := rune(s[i])
		if c == ' ' || c == '\t' || c == '\n' {
			if word != "" {
				strings = append(strings, word)
				word = ""
			} else {
				continue
			}
		} else {
			word += string(s[i])
		}
	}
	if word != "" {
		strings = append(strings, word)
	}
	return strings
}
