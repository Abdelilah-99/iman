package piscine

func NRune(s string, n int) rune {
	tab := []rune(s)
	for i := 0; i < len(tab); i++ {
		if i == n-1 {
			return tab[i]
		}
	}
	return 0
}
