package piscine

func IterativePower(nb int, power int) int {
	r := 1
	if power < 0 {
		return 0
	}
	for i := 1; i <= power; i++ {
		r *= nb
	}
	return r
}
