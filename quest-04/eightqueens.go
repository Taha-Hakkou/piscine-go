package piscine

import "github.com/01-edu/z01"

var (
	n     int = 8
	table [8]int
)

func EightQueens() {
	if n == 0 {
		for j := 0; j < len(table); j++ {
			z01.PrintRune(rune(table[j] + int('0')))
		}
		z01.PrintRune('\n')
		n++
		return
	}
	for i := 1; i <= 8; i++ {
		j := 0
		for j = 0; j < 8-n; j++ {
			if table[j] == i || table[j] == i+n+j-8 || table[j] == 8-n-j+i {
				break
			}
		}
		if j == 8-n {
			table[8-n] = i
			n--
			EightQueens()
		}
	}
	if n < 8 {
		n++
	}
}
