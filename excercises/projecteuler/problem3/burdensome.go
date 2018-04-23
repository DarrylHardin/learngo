package main

import "fmt"

// this would possibly work but the slice and nested loops are too burdensome
//I will upload a better solution

func main() {

	prime := []int{}

	for i := 2; i <= 13195; i++ {
		if i%2 != 0 {
			prime = append(prime, i)
		}
	}
	for i := 2; i <= 13195; i++ {
		for x := 13195; x > 1; x-- {
			for num, key := range prime {
				if i*x == key {
					prime = append(prime[:num], prime[num+1:]...)
					break
				}
			}
		}

	}

	var x int
	for _, key := range prime {
		if 13195%key == 0 {
			if x < key {
				x = key
			}
		}
	}
	fmt.Println(prime)
	fmt.Println(x)
}

//https://projecteuler.net/problem=3
//The prime factors of 13195 are 5, 7, 13 and 29.
//What is the largest prime factor of the number 600851475143 ?
