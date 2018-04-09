package main

import "fmt"

func main() {
	// I was told this code wouldn't work with 5 and 10, I ran the code to find out.
	for i := 1; i <= 100; i++ {
		var f string
		var b string
		if i%5 == 0 {
			f = "Fizz"
		}
		if i%10 == 0 {
			b = "Buzz"
		}

		fmt.Println(i, f+b)
	}
}
