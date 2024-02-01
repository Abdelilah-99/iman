package piscine

func IsUpper(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			if i == len(s)-1 {
				return true
			}
		} else {
			return false
		}
	}
	return false
}
