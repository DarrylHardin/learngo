package main

import (
	"fmt"
)

func checkPalindrome(inputString string) bool {
	l := len(inputString)
	for i := 0; i < (l / 2); i++ {
		if inputString[i] != inputString[l-1-i] {
			return false
		}
	}
	return true
}

func main() {
	x := "aabba"
	y := "lol"
	fmt.Println(checkPalindrome(x))
	fmt.Println(checkPalindrome(y))
}
