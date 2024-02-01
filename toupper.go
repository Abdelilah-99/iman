package piscine

func ToUpper(s string) string {
	str := ""
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			str += string(s[i])
		} else if s[i] >= 'a' && s[i] <= 'z' {
			str += string(s[i] - 32)
		} else {
			str += string(s[i])
		}
	}
	return str
}
