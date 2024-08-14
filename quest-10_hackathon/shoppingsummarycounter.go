package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	cart := make(map[string]int)
	items := []string{}
	word := ""
	for _, r := range str {
		if r == ' ' {
			items = append(items, word)
			word = ""
		} else {
			word += string(r)
		}
	}
	items = append(items, word)
	for _, item := range items {
		if _, ok := cart[item]; ok {
			cart[item]++
		} else {
			cart[item] = 1
		}
	}
	return cart
}
