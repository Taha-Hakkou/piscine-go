package piscine

func MakeRange(min, max int) []int {
	var ints []int
	if max > min {
		ints = make([]int, max-min)
	}
	for i := min; i < max; i++ {
		ints[i-min] = i
	}
	return ints
}
