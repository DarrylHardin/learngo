package main

import "fmt"

//my answer
func centuryFromYear(year int) int {
	century := year / 100
	if year%100 != 0 {
		century++
	}
	return century
}

//callmekit
func centuryFromYear(year int) int {
	if year%100 == 0 {
		return year / 100
	}

	return (year / 100) + 1
}

//bhs on codefights answer

func centuryFromYear(year int) int {
	return 1 + ((year - 1) / 100)
}

func main() {
	fmt.Println(centuryFromYear(2220))
}
