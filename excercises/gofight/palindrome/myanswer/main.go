package main

import (
	"fmt"
	"strings"
)

func checkPalindrome(inputString string) bool {
	f := strings.Split(inputString, "")
	x := len(f)
	b := make([]string, x)
	copy(b[:], f)
	i := len(f)
	for _, y := range f {
		i--
		b[i] = y
	}
	o := strings.Join(f, "")
	r := strings.Join(b, "")

	return o == r
}

func main() {
	x := "aabba"
	y := "lol"
	fmt.Println(checkPalindrome(x))
	fmt.Println(checkPalindrome(y))
}
