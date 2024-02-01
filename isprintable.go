package piscine

func IsPrintable(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] >= 32 {
			if i == len(s)-1 {
				return true
			}
		} else {
			return false
		}
	}
	return false
}
