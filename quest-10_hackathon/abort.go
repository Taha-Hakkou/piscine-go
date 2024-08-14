package piscine

func Abort(a, b, c, d, e int) int {
	args := []int{a, b, c, d, e}
	for i := 0; i <= len(args)/2; i++ {
		for j, n := range args[i+1:] {
			if n < args[i] {
				args[i], args[i+j+1] = args[i+j+1], args[i]
			}
		}
	}
	return args[len(args)/2]
}
