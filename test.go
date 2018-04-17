package main

import "fmt"

func main() {
	valslice := []int{1, 2, 3}
	for i := range valslice {
		val := valslice[i]
		go func() {
			fmt.Println(val)
		}()
		fmt.Println(val)
	}
}
