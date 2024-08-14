package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		PrintPositiveNbr(n / (-10))
		runePlusZero(-n % 10)
	} else {
		PrintPositiveNbr(n)
	}
}

func runePlusZero(n int) {
	c := '0'
	isPositive := true
	if n < 0 {
		n *= (-1)
		isPositive = false
	}
	for i := 0; i < n; i++ {
		if isPositive {
			c++
		} else {
			c--
		}
	}
	z01.PrintRune(c)
}

func PrintPositiveNbr(n int) {
	if n < 10 {
		runePlusZero(n)
		return
	}
	PrintPositiveNbr(n / 10)
	runePlusZero(n % 10)
}

func main() {
	points := &point{}

	setPoint(points)

	printStr("x = ")
	PrintNbr(points.x)
	printStr(", y = ")
	PrintNbr(points.y)
	z01.PrintRune('\n')
}
