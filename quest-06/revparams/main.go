package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	myArgs := os.Args[1:]
	for a := len(myArgs) - 1; a >= 0; a-- {
		runes := []rune(myArgs[a])
		for i := 0; i < len(runes); i++ {
			z01.PrintRune(runes[i])
		}
		z01.PrintRune('\n')
	}
}
