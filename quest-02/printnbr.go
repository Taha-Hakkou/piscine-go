package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		PrintPositiveNbr(n / (-10))
		z01.PrintRune(rune(int('0') - n%10))
	} else {
		PrintPositiveNbr(n)
	}
}

func PrintPositiveNbr(n int) {
	if n < 10 {
		z01.PrintRune(rune(int('0') + n))
		return
	}
	PrintPositiveNbr(n / 10)
	z01.PrintRune(rune((n % 10) + int('0')))
}
