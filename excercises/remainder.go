package main

import (
	"fmt"
)

// set global var
var Num3 int

// DRY: to give message and repeat prompt for input
func check(num int) {
	if num == 0 {
		fmt.Println("You cannot divide a number by 0")
		fmt.Println("Please enter a new number")
	} else {
		Num3 = 1
	}
}

//print remainder between two user inputs
func main() {
	var num1 int
	var num2 int

	for {
		fmt.Print("First Number: ")
		fmt.Scanf("%d\n", &num1)
		check(num1)
		if Num3 == 1 {
			Num3 = 0
			break
		}
	}

	for {
		fmt.Print("Second Number: ")
		fmt.Scanf("%d\n", &num2)
		check(num2)
		if Num3 == 1 {
			break
		}
	}

	if num1 > num2 {
		fmt.Println("Remainder:", num1%num2)
	} else {
		fmt.Println("Remainder:", num2%num1)
	}
}
