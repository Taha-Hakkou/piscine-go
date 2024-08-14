package piscine

func AppendRange(min, max int) []int {
	var ints []int
	for i := min; i < max; i++ {
		ints = append(ints, i)
	}
	return ints
}
