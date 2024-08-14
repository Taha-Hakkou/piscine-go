package piscine

func DescendAppendRange(max, min int) []int {
	ints := []int{}
	for i := max; i > min; i-- {
		ints = append(ints, i)
	}
	return ints
}
