package piscine

func Index(s string, toFind string) int {
	if s == "" || toFind == "" {
		return 0
	}
	j := 0
	for i := 0; i < len(s); i++ {
		if s[i] == toFind[j] {
			j++
			if j == len(toFind) {
				return i - j + 1
			}
		}
	}
	return -1
}
