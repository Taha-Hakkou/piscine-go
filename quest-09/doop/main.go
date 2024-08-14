package main

import "os"

var bytes []byte

func Atoi(s string) int {
	n := 0
	isNegative := false
	for i, c := range s {
		if i == 0 && c == '-' {
			isNegative = true
			continue
		}
		if c < '0' || c > '9' || n*10+int(c)-int('0') < n {
			return 0
		}
		n = n*10 + int(c) - int('0')
	}
	if isNegative {
		n *= (-1)
	}
	return n
}

func printStr(s string) {
	for _, r := range s {
		bytes = append(bytes, byte(r))
	}
	bytes = append(bytes, byte('\n'))
}

func printNbr(n int) {
	if n < 0 {
		bytes = append(bytes, byte('-'))
		if n < -9 {
			PrintPositiveNbr(n / (-10))
		}
		bytes = append(bytes, byte(int('0')-n%10))
	} else {
		PrintPositiveNbr(n)
	}
	bytes = append(bytes, byte('\n'))
}

func PrintPositiveNbr(n int) {
	if n < 10 {
		bytes = append(bytes, byte(int('0')+n))
		return
	}
	PrintPositiveNbr(n / 10)
	bytes = append(bytes, byte((n%10)+int('0')))
}

func main() {
	if len(os.Args) != 4 {
		return
	}
	a := Atoi(os.Args[1])
	if a == 0 {
		for _, r := range os.Args[1] {
			if r != '0' {
				return
			}
		}
	}
	b := Atoi(os.Args[3])
	if b == 0 {
		for _, r := range os.Args[3] {
			if r != '0' {
				return
			}
		}
	}
	switch op := os.Args[2]; op {
	case "+":
		if (a > 0 && b > 0 && a+b <= 0) || (a < 0 && b < 0 && a+b >= 0) {
			return
		}
		printNbr(a + b)
	case "-":
		if (a > 0 && b < 0 && a-b <= 0) || (a < 0 && b > 0 && a-b >= 0) {
			return
		}
		printNbr(a - b)
	case "/":
		if b == 0 {
			printStr("No division by 0")
			os.Stdout.Write(bytes)
			return
		}
		printNbr(a / b)
	case "*":
		absA, absB := a, b
		if a < 0 {
			absA = -a
		}
		if b < 0 {
			absB = -b
		}
		min, max := absA, absB
		if min > max {
			min, max = max, min
		}
		MaxInt := int(^uint(0) >> 1)
		if min != 0 && max != 0 && MaxInt/min < max {
			return
		}
		printNbr(a * b)
	case "%":
		if b == 0 {
			printStr("No modulo by 0")
			os.Stdout.Write(bytes)
			return
		}
		printNbr(a % b)
	default:
		return
	}
	os.Stdout.Write(bytes)
}
