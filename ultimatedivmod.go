package piscine

func UltimateDivMod(a *int, b *int) {
	t := *a / *b
	*b = *a % *b
	*a = t
}
