package piscine

func Capitalize(s string) string {
	str := ""
	isNewWord := true
	var c rune
	for i := 0; i < len(s); i++ {
		c = rune(s[i])
		if c >= 'A' && c <= 'Z' && !isNewWord {
			str += string(c + 'a' - 'A')
		} else if c >= 'a' && c <= 'z' && isNewWord {
			str += string(c + 'A' - 'a')
			isNewWord = false
		} else {
			if (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') {
				isNewWord = false
			} else if c < 'a' || c > 'z' {
				isNewWord = true
			}
			str += string(c)
		}
	}
	return str
}
