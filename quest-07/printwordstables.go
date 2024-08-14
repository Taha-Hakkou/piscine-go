package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, s := range a {
		for j := 0; j < len(s); j++ {
			z01.PrintRune(rune(s[j]))
		}
		z01.PrintRune('\n')
	}
}
