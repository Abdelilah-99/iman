package piscine

func ReverseMenuIndex(menu []string) []string {
	stri := make([]string, len(menu))
	for i := 0; i < len(menu); i++ {
		stri[i] = menu[len(menu)-i-1]
	}
	return stri
}
