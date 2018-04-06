package main

import "fmt"

func main() {
	// will not work with numbers such as 5 and 10
	for i := 1; i <= 100; i++ {
		var f = "Fizz"
		var b = "Buzz"
		switch {
		case i%15 == 0:
			fmt.Println(i, "--", f+b)
		case i%3 == 0:
			fmt.Println(i, "--", f)
		case i%5 == 0:
			fmt.Println(i, "--", b)
		default:
			fmt.Println(i)
		}
	}
}
