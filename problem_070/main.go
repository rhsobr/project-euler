package main

import (
	"fmt"
	"math"
	"project_euler/utils"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func arePermutation(n1, n2 int) bool {
	str1 := strconv.Itoa(n1)
	str2 := strconv.Itoa(n2)

	if len(str1) != len(str2) || len(str1) == 1 {
		return false
	}

	str1X := strings.Split(str1, "")
	str2X := strings.Split(str2, "")

	slices.Sort(str1X)
	slices.Sort(str2X)

	ordered2 := strings.Join(str1X, "")
	ordered1 := strings.Join(str2X, "")

	return ordered1 == ordered2

}

func main() {
	limit := int(math.Pow(10, 6))

	minimum := float64(limit)
	selectedN := 0
	selectPhi := 0

	phis := map[int]int{}

	for n := limit - 1; n >= 3; n -= 2 {
		// they will have many repetitive 2, 3 or 5, then lower Phi results
		if n%3 == 0 || n%5 == 0 {
			continue
		}

		phis[n] = utils.CountNumberOfRelativePrimes(n)

		reason := float64(n) / float64(phis[n])

		if reason < minimum {
			if arePermutation(n, phis[n]) {
				minimum = reason
				selectPhi = phis[n]
				selectedN = n

				fmt.Println("new selecteds", selectedN, selectPhi, minimum)
			}
		}
	}

	fmt.Println("final", selectedN, selectPhi, minimum)
}
