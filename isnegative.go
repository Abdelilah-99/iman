package piscine

import "github.com/01-edu/z01"

func IsNegative(nb int) {
	if nb >= 0 {
		z01.PrintRune(rune('F'))
		z01.PrintRune('\n')
	} else {
		z01.PrintRune(rune('T'))
		z01.PrintRune('\n')
	}
}
