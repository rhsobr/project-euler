package main

import (
	"fmt"
	"math"
	"project_euler/primes"
)

type Fraction struct {
	Numerator   int
	Denominator int
	Ratio       float64
	Stopper     float64
}

func getOrderedFractions(d int) Fraction {
	denominator := d

	stopper := float64(3) / float64(7)
	maxNumerator := 0
	maxDenominator := 1
	ratio := float64(maxNumerator) / float64(maxDenominator)

	for denominator > 1 {
		numerator := int(math.Floor(float64(denominator) * ratio))

		numeratorLimit := int(math.Ceil(float64(denominator) * stopper))

		for numerator <= numeratorLimit {
			result := float64(numerator) / float64(denominator)

			if result > stopper {
				break
			}

			if ratio < result && result < stopper {
				if primes.AreRelativePrimes(denominator, numerator) {

					maxNumerator = numerator
					maxDenominator = denominator
					ratio = float64(maxNumerator) / float64(maxDenominator)

					fmt.Println(Fraction{maxNumerator, maxDenominator, ratio, stopper})
				}
			}

			numerator++
		}

		denominator--
	}

	return Fraction{maxNumerator, maxDenominator, ratio, stopper}
}

func main() {
	fmt.Println(getOrderedFractions(5))
}
