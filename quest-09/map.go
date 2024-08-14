package piscine

func Map(f func(int) bool, a []int) []bool {
	values := []bool{}
	for _, n := range a {
		values = append(values, f(n))
	}
	return values
}
