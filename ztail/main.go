package main

import (
	"os"

	"github.com/01-edu/z01"
)

func toNumber(s string) int {
	n := 0
	for _, r := range s {
		n *= 10
		n += int(rune(r) - '0')
	}
	return n
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func main() {
	var filenames []string
	length := 0
	if os.Args[1] == "-c" {
		filenames = os.Args[3:]
		length = toNumber(os.Args[2])
	} else {
		filenames = os.Args[1:]
	}
	for i, filename := range filenames {
		bytes, err := os.ReadFile(filename)
		if err != nil {
			printStr(err.Error() + "\n")
			os.Exit(1)
		} else if i != 0 {
			printStr("\n")
		}
		if len(filenames) != 1 {
			printStr("==> " + filename + " <==\n")
		}
		printStr(string(bytes[len(bytes)-length:]))
	}
}
