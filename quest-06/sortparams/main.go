package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i := 1; i < len(os.Args)-1; i++ {
		for j := 1; j < len(os.Args)-i; j++ {
			if os.Args[j] > os.Args[j+1] {
				tmp := os.Args[j]
				os.Args[j] = os.Args[j+1]
				os.Args[j+1] = tmp
			}
		}
	}
	for _, v := range os.Args[1:] {
		runes := []rune(v)
		for i := 0; i < len(runes); i++ {
			z01.PrintRune(runes[i])
		}
		z01.PrintRune('\n')
	}
}
