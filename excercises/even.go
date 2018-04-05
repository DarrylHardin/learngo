package main

import (
	"fmt"
)

// print even
func main() {
	for i := 1; i <= 100; i++ {
		answer := i % 2
		switch answer {
		case 0:
			fmt.Println(i, "even")
		}
	}
}
