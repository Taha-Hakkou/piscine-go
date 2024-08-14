package main

import "github.com/01-edu/z01"

func main() {
	for c := 'a'; c <= 'z'; c++ {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
