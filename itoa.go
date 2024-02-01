package main

import (
	"fmt"
)

func main() {
	fmt.Println(Itoa(12345))
	fmt.Println(Itoa(-9223372036854775808))
	fmt.Println(Itoa(-1234))
	fmt.Println(Itoa(987654321))
}

func Itoa(n int) string {
	str := ""
	s := 0
	if n == -9223372036854775808 {
		str += "8"
		s = 1
		n = 922337203685477580
	}
	if n == 0 {
		str = "0"
	}
	if n < 0 {
		n *= -1
		s = 1
	}
	for n > 0 {
		str = string((n%10)+48) + str
		n /= 10
	}
	if s == 1 {
		return "-" + str
	}
	return str
}
