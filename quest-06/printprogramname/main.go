package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	programName := os.Args[0]
	if programName[0:2] == "./" {
		programName = programName[2:]
	}
	runes := []rune(programName)
	for i := 0; i < len(runes); i++ {
		z01.PrintRune(runes[i])
	}
	z01.PrintRune('\n')
}
