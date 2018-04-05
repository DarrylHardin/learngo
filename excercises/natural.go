package main

import (
	"fmt"
)

func main() {
	var sum int
	// var fizz string
	// var buzz string
	var a []int //sets slice
	//loops through 1 to 999
	for i := 1; i < 1000; i++ {
		if i%3 == 0 {
			a = append(a, i)
			// fizz = "fizz"
			// fmt.Println(i)
		} else if i%5 == 0 { // set to else if so int that are divisible by 3 and 5 aren't both counted
			a = append(a, i)
			// buzz = "buzz"
			// fmt.Println(i)
		}

		// fmt.Println(i, fizz, buzz)

		// fizz = ""
		// buzz = ""
	}
	//loops through a and sets sum to the total of the ints in the slice
	for i := range a {
		sum += a[i]
	}
	fmt.Println(sum)
}
