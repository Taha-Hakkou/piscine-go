package main

import "os"

func main() {
	code := os.Args[1]
	const size = 2048
	interpreter := make([]byte, size)
	index := 0
	for i := 0; i < len(code); i++ {
		switch op := rune(code[i]); op {
		case '+':
			interpreter[index]++
		case '-':
			interpreter[index]--
		case '>':
			if index < size-1 {
				index++
			}
		case '<':
			if index > 0 {
				index--
			}
		case '[':
			if interpreter[index] == 0 {
				count := 0
				i++
				for count != 0 || rune(code[i]) != ']' {
					if rune(code[i]) == '[' {
						count++
					} else if rune(code[i]) == ']' {
						count--
					}
					i++
				}
			}
		case ']':
			if interpreter[index] != 0 {
				count := 0
				i--
				for count != 0 || rune(code[i]) != '[' {
					if rune(code[i]) == ']' {
						count++
					} else if rune(code[i]) == '[' {
						count--
					}
					i--
				}
			}
		case '.':
			os.Stdout.Write(interpreter[index:index+1])
		default:
			continue
		}
	}
}
