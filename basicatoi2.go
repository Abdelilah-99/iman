package piscine

func BasicAtoi2(s string) int {
	r := 0
	// str := []byte(s)
	for i := 0; i <= strlen2(s)-1; i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0
		}
		r = (r * 10) + int(s[i]-'0')
	}
	return r
}

func strlen2(s string) int {
	len := 0
	for range s {
		len++
	}
	return len
}
