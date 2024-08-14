package main

import (
	"fmt"
	"piscine"
)

func main() {
	a1 := []int{0, 1, 2, 3, 4, 5}
	a2 := []int{0, 2, 1, 3}

	result1 := piscine.IsSorted(f, a1)
	result2 := piscine.IsSorted(f, a2)

	fmt.Println(result1)
	fmt.Println(result2)

	fmt.Println(piscine.IsSorted(f, []int{559580, 391756, 312015, -21866, -144074, -515500, -643274, -951476}))
}

func f(a, b int) int {
	return a - b
}
