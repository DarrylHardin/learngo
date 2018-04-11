package main

import "fmt"

func main() {
	x, y, z := 3, 24, 25
	var total int
	var high int
	for {
		if y == 3 {
			break
		}
		total = x * y
		fmt.Println(x, y)
		fmt.Println(total)
		if total == z {
			if total%2 == 0 {

				switch {
				case x != y:
					x++
				default:
					x = 3
					y--
				}
				continue

			} else if total%5 == 0 {
				high = total
				x = 3
				y--
			}

		} else {
			switch {
			case x != y:
				x++
			default:
				x = 3
				y--
			}
			continue
		}
	}

	fmt.Println(high)

}
