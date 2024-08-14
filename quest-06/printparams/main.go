package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for _, v := range os.Args[1:] {
		runes := []rune(v)
		for i := 0; i < len(runes); i++ {
			z01.PrintRune(runes[i])
		}
		z01.PrintRune('\n')
	}
}
