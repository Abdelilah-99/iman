package piscine

func IsNumeric(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			if i == len(s)-1 {
				return true
			}
		} else {
			return false
		}
	}
	return false
}
