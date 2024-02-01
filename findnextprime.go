package piscine

func FindNextPrime(nb int) int {
	if isprime(nb) {
		return nb
	}
	for i := nb; i <= 9223372036854775807; i++ {
		if isprime(i) {
			return i
		}
	}
	return 2
}

func isprime(nb int) bool {
	if nb == 1 || nb <= 0 {
		return false
	}
	index := 0
	for i := 1; i <= nb/i; i++ {
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
