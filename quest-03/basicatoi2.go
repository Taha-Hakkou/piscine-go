package piscine

func BasicAtoi2(s string) int {
	n := 0
	for i := 0; i < len(s); i++ {
		c := rune(s[i])
		if c < '0' || c > '9' {
			return 0
		}
		n = n*10 + int(c) - int('0')
	}
	return n
}
