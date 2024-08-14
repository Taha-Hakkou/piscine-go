package piscine

func ConcatParams(args []string) string {
	var s string
	if len(args) > 0 {
		s = args[0]
	}
	for i := 1; i < len(args); i++ {
		s += "\n" + args[i]
	}
	return s
}
