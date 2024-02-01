package main

import "github.com/01-edu/z01"

func main() {
	str := "hello word"
	for _, i := range str {
		z01.PrintRune(rune(i))
	}
	z01.PrintRune('\n')
}
