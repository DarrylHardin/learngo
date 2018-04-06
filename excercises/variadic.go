package main

import "fmt"

// find the highest number in a slice
func main() {
	data := []int{10, 20, 30, 5, 22, 99, 60, 30}
	n := highest(data...)
	fmt.Println(n)
}

func highest(hn ...int) int {
	var h int
	for _, x := range hn {
		if x > h {
			h = x
		}
	}
	return h
}
