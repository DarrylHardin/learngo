package main

import (
	"fmt"
)

// simple code to sending struct through func
type name struct {
	first string
	last  string
}

func who(f, l string) (string, string) {
	return f, l
}

func passStruct(fullname name) (string, string) {
	return fullname.first, fullname.last
}

func main() {
	me := name{"Darryl", "Hardin"}
	//this one takes the parameters of the name struct as separate values
	fmt.Println(who(me.first, me.last))
	//this one takes the complete struct
	fmt.Println(passStruct(me))
}
