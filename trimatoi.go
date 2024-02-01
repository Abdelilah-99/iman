package piscine

func TrimAtoi(s string) int {
	sign := 1
	str := 0
	f := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			str = (str * 10) + int(s[i]-48)
			f = 1
		}
		if f == 0 && s[i] == '-' {
			sign = -1
		}
	}
	return str * sign
}
