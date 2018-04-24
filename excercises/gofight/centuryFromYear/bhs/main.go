package main

import "fmt"

func centuryFromYear(year int) int {
	return 1 + ((year - 1) / 100)
}

func main() {
	fmt.Println(centuryFromYear(2220))
}
