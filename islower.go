package piscine

func IsLower(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			if i == len(s)-1 {
				return true
			}
		} else {
			return false
		}
	}
	return false
}
