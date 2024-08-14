package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	s := ""
	if len(args) > 0 {
		s = args[0]
		for _, arg := range args[1:] {
			s += " " + arg
		}
	}
	vowels := []rune{'a', 'A', 'e', 'E', 'i', 'I', 'o', 'O', 'u', 'U'}
	isVowel := false
	i, j := 0, len(s)-1
	for i < j {
		for i < j {
			for _, vowel := range vowels {
				if rune(s[i]) == vowel {
					isVowel = true
					break
				}
			}
			if isVowel {
				isVowel = false
				break
			}
			i++
		}
		for i < j {
			for _, vowel := range vowels {
				if rune(s[j]) == vowel {
					isVowel = true
					break
				}
			}
			if isVowel {
				isVowel = false
				break
			}
			j--
		}
		if i != j {
			s = s[:i] + string(s[j]) + s[i+1:j] + string(s[i]) + s[j+1:]
		}
		i++
		j--
	}
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
