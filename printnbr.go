package piscine

import (
	"fmt"
)

func PrintNbr(n int) {
	if n == 0 {
		fmt.Print('0')
		return
	}
	if n == -9223372036854775808 {
		fmt.Print('-')
		fmt.Print('9')
		n = 223372036854775808
	}
	if n < 0 {
		fmt.Print('-')
		n = n * (-1)
	}
	rec(n)
}

func rec(n int) {
	if n >= 10 {
		rec(n / 10)
		digit := n % 10
		fmt.Print(digit)
	}
	if n < 10 {
		digit := n % 10
		fmt.Print(digit)
	}
}
