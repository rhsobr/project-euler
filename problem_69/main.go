package main

import (
	"fmt"
	"project_euler/primes"
)

func countNumberOfRelativePrimes(number int, counterStopAt float64) int {
	position := 2

	counter := 1

	if primes.IsPrime(number) {
		return number - 1
	}

	for position < number {
		if primes.AreRelativePrimes(position, number) {
			counter++

			if counterStopAt > 0 && float64(number)/float64(counter) < 5 {
				return counter
			}
		}
		position++
	}

	return counter
}

func getMaximunPositionOfPhi(number int) int {
	position := number

	maxPhi := 0.0
	maxPosition := 0

	for position > 0 {

		relativePrimes := countNumberOfRelativePrimes(position, maxPhi)

		result := float64(position) / float64(relativePrimes)

		if maxPhi < result {
			maxPhi = result
			maxPosition = position

			fmt.Print(position)
			fmt.Print(" ")
			fmt.Print(relativePrimes)
			fmt.Print(" ")
			fmt.Println(result)
		}

		position -= 10
	}

	return maxPosition
}

func main() {
	fmt.Println(getMaximunPositionOfPhi(1000000))
}
