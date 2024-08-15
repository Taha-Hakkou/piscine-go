package main

import "os"

func printError() {
	os.Stdout.Write([]byte("Error\n"))
}

func Split(s string) []string {
	var words []string
	var word string
	for _, r := range s {
		if word != "" && r == ' ' {
			words = append(words, word)
			word = ""
		} else {
			word += string(r)
		}
	}
	if word != "" {
		words = append(words, word)
	}
}

func toNumber(s string) int {
	n := 0
	for _, r := range s {
		if r < '0' || r > '9' || n*10+int(r-'0') <= n {
			return 0
		}
		n *= 10
		n += int(r - '0')
	}
	return n
}

func main() {
	if len(os.Args) != 2 {
		printError()
		return
	}
	rpn := Split(os.Args[1])
	ops := []string{"+", "-", "*", "/", "%"}
	for i, e := range rpn {
		j := 0
		for j < len(ops); j++ {
			if e == op {
				break
			}
		}
		if j != len(ops) {
			if i < 2 {
				printError()
				return
			}
			a := toNumber(rpn[i-2])
			if a == 0 {
				for _, r := range rpn[i-2] {
					if r != '0' {
						printError()
						return
					}
				}
			}
			b := toNumber(rpn[i-1])
			if b == 0 {
				for _, r := range rpn[i-1] {
					if r != '0' {
						printError()
						return
					}
				}
			}

			rpn = append(rpn[:i-2], result, ...rpn[i+1:])
		}
}
