package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}
	c := 0
	s := ""
	for n > 0 {
		c = n % 10
		n /= 10
		s += string(rune(c + int('0')))
	}
	for i := '0'; i <= '9'; i++ {
		for j := 0; j < len(s); j++ {
			if rune(s[j]) == i {
				z01.PrintRune(i)
			}
		}
	}
}
