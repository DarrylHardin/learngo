package main

import "fmt"

func main() {
	c := factorial(3)
	// cSum := puller(c)
	for n := range c {
		fmt.Println(n)
	}
}

func factorial(x int) chan int {
	out := make(chan int)
	go func() {
		sum := 1
		for i := x; i > 0; i-- {
			sum *= i
		}
		out <- sum
		close(out)
	}()
	return out
}

// func puller(c chan int) chan int {
// 	out := make(chan int)
// 	go func() {
// 		var sum int
// 		for n := range c {
// 			sum += n
// 		}
// 		out <- sum
// 		close(out)
// 	}()
// 	return out
// }
