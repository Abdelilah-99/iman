package piscine

func JumpOver(str string) string {
	stri := ""
	for i := 0; i < len(str); i++ {
		if i == 2 {
			stri += string(str[i])
			str = str[i+1:]
			i = 0
		}
	}
	return stri + "\n"
}
