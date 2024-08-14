package piscine

func StringToIntSlice(str string) []int {
	var ints []int
	for _, r := range str {
		ints = append(ints, int(r))
	}
	return ints
}
