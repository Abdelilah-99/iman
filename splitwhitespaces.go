package piscine

func SplitWhiteSpaces(s string) []string {
	var arr []string
	str := ""
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			str += string(s[i])
		} else if str != "" {
			arr = append(arr, str)
			str = ""
		}
	}
	if str != "" {
		arr = append(arr, str)
		str = ""
	}
	return arr
}
