package piscine

func ToLower(s string) string {
	runes := []rune(s)
	str := ""
	for i := 0; i < len(runes); i++ {
		if runes[i] >= 'A' && runes[i] <= 'Z' {
			str += string(runes[i] + 'a' - 'A')
		} else {
			str += string(runes[i])
		}
	}
	return str
}
