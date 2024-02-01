package piscine

func IsPrime(nb int) bool {
	if nb == 1 || nb <= 0 {
		return false
	}
	index := 0
	for i := 2; i <= nb; i++ {
		if nb%i == 0 {
			index++
			if index >= 2 {
				return false
			}
		}
	}
	if index == 1 {
		return true
	}
	return false
}
