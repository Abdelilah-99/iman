package piscine

func RecursivePower(nb int, power int) int {
	r := 1
	if power < 0 {
		return 0
	}
	if power >= 1 {
		r = nb * RecursivePower(nb, power-1)
	}
	return r
}
