package piscine

func Compact(ptr *[]string) int {
	s := 0
	var arr []string
	for _, ad := range *ptr {
		if ad != "" {
			arr = append(arr, ad)
			s++
		}
	}
	*ptr = arr
	return s
}
