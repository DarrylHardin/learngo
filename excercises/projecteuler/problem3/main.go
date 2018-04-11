package main

import "fmt"

func main() {

	check := 13195
	var highest int
	for i := 2; i < check; i++ {
		switch {
		case i%2 == 0 && i != 2:
			continue
		case i%3 == 0 && i != 3:
			continue
		case i%4 == 0 && i != 4:
			continue
		case i%5 == 0 && i != 5:
			continue
		case i%6 == 0 && i != 6:
			continue
		case i%7 == 0 && i != 7:
			continue
		case i%8 == 0 && i != 8:
			continue
		case i%9 == 0 && i != 9:
			continue
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
