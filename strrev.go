package piscine

func StrRev(s string) string {
	str := ""
	for i := strlen(s) - 1; i >= 0; i-- {
		str += string(s[i])
	}
	return str
}

func strlen(s string) int {
	len := 0
	for range s {
		len++
	}
	return len
}
