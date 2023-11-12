package main

import (
	"fmt"
	"math"
	"project_euler/utils"
	"reflect"

	"golang.org/x/exp/slices"
)

func arePermutation(n1, n2 int) bool {
	if n1 < 10 || n2 < 10 {
		return false
	}

	n1Digits := utils.GetDigits(n1)
	n2Digits := utils.GetDigits(n2)

	if len(n1Digits) != len(n2Digits) {
		return false
	}

	slices.Sort(n1Digits)
	slices.Sort(n2Digits)

	return reflect.DeepEqual(n1Digits, n2Digits)
}

func main() {
	limit := int(math.Pow(10, 7))

	minimum := float64(limit)
	selectedN := 0
	selectPhi := 0

	for n := limit - 1; n >= 3; n -= 2 {
		// they will have many repetitive 2, 3 or 5, then lower Phi results
		if n%3 == 0 || n%5 == 0 {
			continue
		}

		phi := utils.CountNumberOfRelativePrimes(n)

		reason := float64(n) / float64(phi)

		if reason < minimum {
			if arePermutation(n, phi) {
				minimum = reason
				selectPhi = phi
				selectedN = n

				fmt.Println("new selecteds", selectedN, selectPhi, minimum)
			}
		}
	}

	fmt.Println("final", selectedN, selectPhi, minimum)
}
