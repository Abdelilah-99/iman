package piscine

func BasicAtoi(s string) int {
	r := 0
	// str := []byte(s)
	for i := 0; i <= strlen1(s)-1; i++ {
		r = (r * 10) + int(s[i]-'0')
	}
	return r
}

func strlen1(s string) int {
	len := 0
	for range s {
		len++
	}
	return len
}
