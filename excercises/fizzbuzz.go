package main

import (
	"fmt"
)

func main() {
	var fizz string
	var buzz string
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fizz = "fizz"
		}
		if i%5 == 0 {
			buzz = "buzz"
		}

		fmt.Println(i, fizz, buzz)

		fizz = ""
		buzz = ""
	}
}
