package main

import (
	"fmt"
)

func main() {

	half := func(x int) (int, bool) { //pass in one int and return int, bool
		return x / 2, x%2 == 0
	}

	fmt.Println(half(10))
}
