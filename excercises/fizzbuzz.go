package main

import "fmt"

func main() {
	// will not work with numbers such as 5 and 10
	for i := 1; i <= 100; i++ {
		var fizz string
		var buzz string
		if i%3 == 0 {
			fizz = "Fizz"
		}
		if i%5 == 0 {
			buzz = "Buzz"
		}

		fmt.Println(i, fizz+buzz)
	}
}
