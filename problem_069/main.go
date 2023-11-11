package main

import (
	"fmt"
	"project_euler/utils"
)

func getMaximunPositionOfPhi(number int) int {
	position := number

	maxPhi := 0.0
	maxPosition := 0

	if position%2 == 1 {
		position = position - 1
	}

	for position > 0 {
		relativePrimes := utils.CountNumberOfRelativePrimes(position)

		result := float64(position) / float64(relativePrimes)

		if maxPhi <= result {
			maxPhi = result
			maxPosition = position

			fmt.Println(position, relativePrimes, result)
		}

		position -= 10
	}

	return maxPosition
}

func main() {
	fmt.Println(getMaximunPositionOfPhi(1_000_000))
}
