package main

import "fmt"

type robot struct {
	weapon string
	heat   int
	status bool
}

func main() {

	r1 := robot{"Laser", 20, true}
	r2 := robot{"Gatling", 50, false}

	fmt.Println(r1.weapon, r1.heat, r1.status)
	fmt.Println(r2.weapon, r2.heat, r2.status)
}
