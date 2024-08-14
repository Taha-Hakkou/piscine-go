package piscine

func ToUpper(s string) string {
	runes := []rune(s)
	str := ""
	for i := 0; i < len(runes); i++ {
		if runes[i] >= 'a' && runes[i] <= 'z' {
			str += string(runes[i] - 'a' + 'A')
		} else {
			str += string(runes[i])
		}
	}
	return str
}
