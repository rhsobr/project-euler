package main

import (
	"fmt"
	"project_euler/utils"
)

func main() {
	limit := 1_000_000

	relativePrimes := int64(0)

	for n := limit; n >= 1; n-- {
		relatives := utils.CountNumberOfRelativePrimes(n)
		relativePrimes += int64(relatives)
	}

	fmt.Println("Sum of Relative Primes - 1", relativePrimes-1)
}
