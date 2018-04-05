package main

import (
	"fmt"
)

func half(num int) (x int) {
	var y bool
	y = false
	if a := num; num%2 == 0 {
		x = a / 2
		y = true
	}
	return
}

func main() {

	fmt.Println(half(2))
}
