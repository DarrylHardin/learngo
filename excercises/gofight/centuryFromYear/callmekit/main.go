package main

import "fmt"

//callmekit
func centuryFromYear(year int) int {
	if year%100 == 0 {
		return year / 100
	}

	return (year / 100) + 1
}

func main() {
	fmt.Println(centuryFromYear(2220))
}
