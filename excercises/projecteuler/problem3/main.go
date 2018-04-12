package main

import "fmt"

func main() {
	check := int64(600851475143)
	var highest int64
	for i := check / int64(3); i < check/int64(2); i++ {
		switch {
		case i%2 == 0 && i != 2:
			continue
		case i%3 == 0 && i != 3:
			continue
		case i%5 == 0 && i != 5:
			continue
		case i%7 == 0 && i != 7:
			continue
		case i%11 == 0 && i != 11:
			continue
		case i%13 == 0 && i != 13:
			continue
		// case i%8 == 0 && i != 17:
		// 	continue
		// case i%9 == 0 && i != 19:
		// 	continue
		default:
			if check%i == 0 {
				if i > highest {
					highest = i
				}
			}
		}
	}
	fmt.Println(highest)
}
