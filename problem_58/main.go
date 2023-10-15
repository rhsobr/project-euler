package main

import (
	"fmt"
	"math"
	"project_euler/primes"
)

func getCorners(n int) []int {
	factor := (n + (n - 1))

	bottomRight := int(math.Pow(float64(factor), 2))

	stepper := factor - 1

	bottomLeft := bottomRight - stepper

	topLeft := bottomLeft - stepper

	topRight := topLeft - stepper

	return []int{bottomRight, bottomLeft, topLeft, topRight}

}

func main() {
	primeCorners := 0
	corners := 1

	var lastCorners []int

	for sideLength := 2; ; sideLength++ {
		lastCorners = getCorners(sideLength)

		corners += len(lastCorners)

		for _, corner := range lastCorners[1:] {
			if primes.IsPrime(corner) {
				primeCorners++
			}
		}

		ratio := float64(primeCorners) / float64(corners)

		if ratio < .1 {
			break
		}
	}

	fmt.Println("Answer:", math.Sqrt(float64(lastCorners[0])))
}
