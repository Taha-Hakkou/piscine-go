package piscine

func Atoi(s string) int {
	n := 0
	isNegative := false
	for i := 0; i < len(s); i++ {
		c := rune(s[i])
		if i == 0 && (c == '+' || c == '-') {
			if c == '-' {
				isNegative = true
			}
			continue
		}
		if c < '0' || c > '9' {
			return 0
		}
		n = n*10 + int(c) - int('0')
	}
	if isNegative {
		n *= (-1)
	}
	return n
}
