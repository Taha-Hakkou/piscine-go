package piscine

func Compact(ptr *[]string) int {
	n := 0
	for i, s := range *ptr {
		if s != "" {
			(*ptr)[n], (*ptr)[i] = (*ptr)[i], (*ptr)[n]
			n++
		}
	}
	*ptr = (*ptr)[:n]
	return n
}
