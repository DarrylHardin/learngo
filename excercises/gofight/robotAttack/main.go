package main

import (
	"fmt"
)

func bot(robots []int) int {
	var s int
	var e bool
	var i int
	for i < len(robots) {
		health := robots[i]
		if e == true {
			health -= 2
			if health <= 0 {
				robots = append(robots[:i], robots[i+1:]...)
				if i == len(robots) {
					e = false
					i = 0
				}

				continue
			}
		}
		if health <= 0 {
			robots = append(robots[:i], robots[i+1:]...)
			if i == len(robots) {
				e = false
				i = 0
			}
			continue
		}
		health--
		s++
		if health == 0 {
			robots = append(robots[:i], robots[i+1:]...)
			e = true
			if i == len(robots) {
				e = false
				i = 0
			}
			continue
		} else {
			robots[i] = health
			i++

			if len(robots) == 0 {
				break
			} else if i == len(robots) {
				e = false
				i = 0
				continue
			}
		}
	}
	fmt.Println(robots)
	return s
}
func main() {
	r := []int{1, 2, 2, 3, 1, 3}
	x := []int{0, 4, 2}
	fmt.Println(bot(r))
	fmt.Println(bot(x))
}
