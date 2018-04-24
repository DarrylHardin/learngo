package main

import (
	"fmt"
)

func checkPalindrome(inputString string) bool {
	strlen := len(inputString)
	str := inputString

	if strlen%2 == 0 {
		return str[:strlen/2] == str[strlen/2:]
	}

	return str[:strlen/2] == str[(strlen/2)+1:]
}
func main() {
	x := "aabba"
	y := "lol"
	fmt.Println(checkPalindrome(x))
	fmt.Println(checkPalindrome(y))
}
