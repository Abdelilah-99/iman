package piscine

func StringToIntSlice(str string) []int {
	var resultat []int
	for _, i := range str {
		resultat = append(resultat, int(i))
	}
	return resultat
}
