package piscine

func Split(s, sep string) []string {
	var arr []string
	str := ""
	for i := 0; i < len(s); i++ {
		if i+len(sep) <= len(s) && s[i:i+len(sep)] == sep {
			arr = append(arr, str)
			str = ""
			i += len(sep) - 1
		} else {
			str += string(s[i])
		}
	}
	arr = append(arr, str)
	return arr
}
