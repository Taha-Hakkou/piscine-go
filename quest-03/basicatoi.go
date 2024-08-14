package piscine

func BasicAtoi(s string) int {
	n := 0
	for i := 0; i < len(s); i++ {
		c := rune(s[i])
		n = n*10 + int(c) - int('0')
	}
	return n
}
