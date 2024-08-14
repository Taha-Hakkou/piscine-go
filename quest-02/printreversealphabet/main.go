package main

import "github.com/01-edu/z01"

func main() {
	for c := 'z'; c >= 'a'; c-- {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
