package main

import (
	"fmt"
)

func main() {
	var x int
	var p int
	for i := 900; i < 1000; i++ {
		for y := 999; y > 899; y-- {
			if (i*y)%11 == 0 {
				x = i * y
				new_int := 0
				for x > 0 {
					remainder := x % 10
					new_int *= 10
					new_int += remainder
					x /= 10
					if new_int == x {
						p = new_int
					}
				}
			}
		}
	}
	fmt.Println(p)
}
