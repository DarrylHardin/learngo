package main

import (
	"fmt"
	"sort"
)

func main() {
	type people []string
	studyGroup := people{"Dan", "Aaron", "Dave", "Zach", "Eric"}
	a := []int{2, 6, 8, 2, 1, 4}
	sort.Ints(a)
	fmt.Println(a)
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	fmt.Println(a)
	fmt.Println(studyGroup)
}
