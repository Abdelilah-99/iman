package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if len(base) < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	for i := 0; i < len(base); i++ {
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] || base[i] == '+' || base[i] == '-' {
				z01.PrintRune('N')
				z01.PrintRune('V')
				return
			}
		}
	}
	if nbr == -9223372036854775808 {
		z01.PrintRune('-')
		PrintNbrBase(9, base)
		PrintNbrBase(223372036854775808, base)
		return
	}
	if nbr < 0 {
		z01.PrintRune('-')
		nbr *= -1
	}
	res := ""
	for nbr != 0 {
		res = string(base[nbr%len(base)]) + res
		nbr /= len(base)
	}
	printstr(string(res))
}

func printstr(s string) string {
	s1 := []rune(s)
	for i := 0; i < len(s); i++ {
		z01.PrintRune(s1[i])
	}
	return string(s1)
}
