package main

import "os"

func isCorrectlyBracketed(s string) bool {
	openers := []rune{'(', '[', '{'}
	closers := []rune{')', ']', '}'}
	for start := 0; start < len(s); start++ {
		r := rune(s[start])
		for i := range closers {
			if r == closers[i] {
				return false
			}
		}
		for i := range openers {
			if r == openers[i] {
				count := 0
				end := start + 1
				if end >= len(s)-1 {
					return false
				}
				rx := rune(s[end])
				for count != 0 || rx != closers[i] {
					if rx == openers[i] {
						count++
					} else if rx == closers[i] {
						count--
					}
					if end < len(s)-1 {
						end++
					} else {
						return false
					}
					rx = rune(s[end])
				}
				return isCorrectlyBracketed(s[start+1:end]) && isCorrectlyBracketed(s[end+1:])
			}
		}
	}
	return true
}

func main() {
	for _, arg := range os.Args[1:] {
		if isCorrectlyBracketed(arg) {
			os.Stdout.Write([]byte("OK\n"))
		} else {
			os.Stdout.Write([]byte("Error\n"))
		}
	}
}
