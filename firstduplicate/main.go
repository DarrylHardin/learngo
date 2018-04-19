package main

import "fmt"

// func firstDuplicate(a []int) int {

// 	x := len(a) - 1
// 	var y bool
// 	for num, key := range a {
// 		exist := key
// 		for i := num + 1; i < len(a); i++ {
// 			if exist == a[i] {
// 				if x > i {
// 					x = i
// 					y = true
// 				} else if len(a) <= 3 {
// 					x = i
// 					y = true
// 				}
// 			}
// 		}
// 	}
// 	if y == true {
// 		return a[x]
// 	} else {
// 		return -1
// 	}

// }

// func firstDuplicate(a []int) int {
//     var num int
//     low := len(a)
//     var x bool
//     c := make(chan int)
//     go func() {
//         for key,val := range a {
//             if num == val {
//                 if low > key {
//                     low = key
//                     x = true
//                 }
//             }
//             num = val
//         }
//     }

//     for i := range
//     if x == true{
//         return a[low]
//     } else {
//         return -1
//     }

// }

// func firstDuplicate(a []int) int {
// 	seen := []int{}
// 	for _, fVal := range a {
// 		for _, val := range seen {
// 			if val == fVal {
// 				return val
// 			}
// 		}
// 		seen = append(seen, fVal)
// 	}
// 	return seen[0]
// }

func firstDuplicate(a []int) int {
	seen := []int{}
	c := make(chan int)
	x := -1
	go func() {
		for _, fVal := range a {
			c <- fVal
		}
		close(c)
	}()

	go func() {
		for fVal := range c {
			for _, val := range seen {
				if val == fVal {
					x = val
				}
			}
			seen = append(seen, fVal)
		}
	}()
	return x
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
