package main

import "fmt"

func firstDuplicate(a []int) int {

	x := len(a) - 1
	var y bool
	for num, key := range a {
		exist := key
		for i := num + 1; i < len(a); i++ {
			if exist == a[i] {
				if x > i {
					x = i
					y = true
				} else if len(a) <= 3 {
					x = i
					y = true
				}
			}
		}
	}
	if y == true {
		return a[x]
	} else {
		return -1
	}

}

func main() {

	a := []int{4, 3, 3, 5, 4}       //4
	b := []int{5, 6, 5, 6}          //5
	c := []int{3, 4, 1, 4, 5, 8, 8} //4
	d := []int{2, 3, 3, 1, 5, 2}    //3
	fmt.Println(firstDuplicate(a))
	fmt.Println(firstDuplicate(b))
	fmt.Println(firstDuplicate(c))
	fmt.Println(firstDuplicate(d))
}
