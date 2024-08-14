package piscine

func ReverseMenuIndex(menu []string) []string {
	reverse_menu := make([]string, len(menu))
	for i := 0; i < len(menu); i++ {
		reverse_menu[i] = menu[len(menu)-i-1]
	}
	return reverse_menu
}
