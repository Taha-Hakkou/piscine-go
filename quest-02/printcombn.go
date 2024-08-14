package piscine

import "github.com/01-edu/z01"

func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	}
	return nb * RecursivePower(nb, power-1)
}

func PrintNbr(nbr int, n int) {
	if nbr < 10 {
		for i := 0; i < n-1; i++ {
			z01.PrintRune('0')
		}
		z01.PrintRune(rune(int('0') + nbr))
		return
	}
	PrintNbr(nbr/10, n-1)
	z01.PrintRune(rune((nbr % 10) + int('0')))
}

func PrintCombN(n int) {
	p := RecursivePower(10, n)
	firstPrint := true
	for i := 0; i < p; i++ {
		j := i
		m := n - 1
		for m > 0 {
			if j%10 <= (j/10)%10 {
				break
			}
			j /= 10
			m--
		}
		if m == 0 {
			if firstPrint {
				firstPrint = false
			} else {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			PrintNbr(i, n)
		}
	}
	z01.PrintRune('\n')
}
