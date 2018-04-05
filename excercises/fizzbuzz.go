package main

import "fmt"

func main() {
	var fizz string
	var buzz string
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fizz = "Fizz"
		}
		if i%5 == 0 {
			buzz = "Buzz"
		}

		fmt.Println(i, fizz+buzz)

		fizz = ""
		buzz = ""
	}
}
