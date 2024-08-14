package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func main() {
	args := os.Args[1:]
	if len(args) > 0 {
		for _, filename := range args {
			data, err := os.ReadFile(filename)
			if err != nil {
				printStr("ERROR: " + err.Error() + "\n")
				os.Exit(1)
			}
			printStr(string(data))
		}
	} else {
		buffer := make([]byte, 1024)
		n, err := os.Stdin.Read(buffer)
		if err != nil {
			return
		}
		printStr(string(buffer[:n]))
	}
}
