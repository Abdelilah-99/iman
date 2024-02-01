package piscine

func Capitalize(s string) string {
	str := tolower(s)
	str1 := ""
	a := true
	for i := 0; i < len(str); i++ {
		if IsAlpha(string(str[i])) {
			if a == false || i == 0 {
				str1 += toupper(string(str[i]))
				a = true
			} else {
				str1 += string(str[i])
			}
		} else {
			str1 += string(str[i])
			a = false
		}
	}
	return str1
}

func tolower(s string) string {
	str := ""
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			str += string(s[i])
		} else if s[i] >= 'A' && s[i] <= 'Z' {
			str += string(s[i] + 32)
		} else {
			str += string(s[i])
		}
	}
	return str
}

func toupper(s string) string {
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

func isalpha(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' || s[i] >= 'A' && s[i] <= 'Z' || s[i] >= '0' && s[i] <= '9' {
			if i == len(s)-1 {
				return true
			}
		} else {
			return false
		}
	}
	return false
}
