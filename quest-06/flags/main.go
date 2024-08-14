package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 || (len(args) == 1 && (args[0] == "--help" || args[0] == "-h")) {
		printHelp()
	} else {
		inOrder := false
		s := args[len(args)-1]
		for i := 0; i < len(args)-1; i++ {
			if args[i] == "--order" || args[i] == "-o" {
				inOrder = true
			} else if len(args[i]) > 8 && args[i][:9] == "--insert=" {
				s += args[i][9:]
			} else if len(args[i]) > 2 && args[i][:3] == "-i=" {
				s += args[i][3:]
			}
		}
		if inOrder {
			str := []rune(s)
			for i := 0; i < len(str)-1; i++ {
				for j := 0; j < len(str)-1-i; j++ {
					if rune(str[j]) > rune(str[j+1]) {
						tmp := str[j]
						str[j] = str[j+1]
						str[j+1] = tmp
					}
				}
			}
			s = ""
			for i := 0; i < len(str); i++ {
				s += string(str[i])
			}
		}
		fmt.Println(s)
	}
}

func printHelp() {
	fmt.Println("--insert")
	fmt.Println("  -i")
	fmt.Println("\t This flag inserts the string into the string passed as argument.")
	fmt.Println("--order")
	fmt.Println("  -o")
	fmt.Println("\t This flag will behave like a boolean, if it is called it will order the argument.")
}
