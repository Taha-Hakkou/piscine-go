package piscine

func TrimAtoi(s string) int {
	n := 0
	isNegative := false
	for i := 0; i < len(s); i++ {
		if rune(s[i]) == '-' && n == 0 {
			isNegative = true
		} else if rune(s[i]) >= '0' && rune(s[i]) <= '9' {
			n = n*10 + int(rune(s[i])) - int('0')
		}
	}
	if isNegative {
		n *= (-1)
	}
	return n
}
