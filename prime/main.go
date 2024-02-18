package main

import (
	"fmt"
)

func main() {
	var MAX int = 10

	primes := make([]int, 0)

	for i := range MAX {
		// check if this number is prime or not
		fmt.Println(i)

		if i == 0 || i == 1 {
			continue
		}

		if i == 2 {
			primes = append(primes, i)
			continue
		}

		if i%2 != 0 {
			primes = append(primes, i)
		}

	}

	fmt.Println(primes)
}
