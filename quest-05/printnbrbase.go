package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if !isValid(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	if nbr < 0 {
		z01.PrintRune('-')
		PrintPositiveNbrValidBase(nbr/(-len(base)), base)
		z01.PrintRune(rune(base[-(nbr % len(base))]))
	} else {
		PrintPositiveNbrValidBase(nbr, base)
	}
}

func PrintPositiveNbrValidBase(nbr int, base string) {
	if nbr < len(base) {
		z01.PrintRune(rune(base[nbr]))
		return
	}
	PrintPositiveNbrValidBase(nbr/len(base), base)
	z01.PrintRune(rune(base[nbr%len(base)]))
}

func isValid(base string) bool {
	if len(base) < 2 {
		return false
	}
	for i := 0; i < len(base); i++ {
		c := rune(base[i])
		if c == '+' || c == '-' {
			return false
		}
		for j := 0; j < i; j++ {
			if c == rune(base[j]) {
				return false
			}
		}
	}
	return true
}
