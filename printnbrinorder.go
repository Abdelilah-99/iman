package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	var arr []int
	if n == 0 {
		arr = append(arr, 0)
	}
	for n > 0 {
		arr = append(arr, n%10)
		n /= 10
	}
	arr1 := sort(arr)
	for i := 0; i < len(arr); i++ {
		z01.PrintRune(rune(arr1[i]) + 48)
	}
}

func sort(arr1 []int) []int {
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr1)-i-1; j++ {
			if arr1[j] > arr1[j+1] {
				t := arr1[j]
				arr1[j] = arr1[j+1]
				arr1[j+1] = t
			}
		}
	}
	return arr1
}
