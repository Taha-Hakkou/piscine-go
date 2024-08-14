package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	var args []string
	toUpper := false
	if len(os.Args) <= 1 {
		return
	}
	if os.Args[1] == "--upper" {
		toUpper = true
		args = os.Args[2:]
	} else {
		args = os.Args[1:]
	}
	for _, s := range args {
		nbr := 0
		i := 0
		for i = 0; i < len(s); i++ {
			c := rune(s[i])
			if c < '0' || c > '9' {
				z01.PrintRune(' ')
				break
			}
			nbr *= 10
			nbr += int(c) - int('0')
		}
		if i == len(s) {
			if nbr > 26 {
				z01.PrintRune(' ')
			} else if toUpper {
				z01.PrintRune(rune(nbr - 1 + int('A')))
			} else {
				z01.PrintRune(rune(nbr - 1 + int('a')))
			}
		}
	}
	z01.PrintRune('\n')
}
