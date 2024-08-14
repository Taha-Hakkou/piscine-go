package piscine

func UltimateDivMod(a *int, b *int) {
	t := *a % *b
	*a = *a / *b
	*b = t
}
