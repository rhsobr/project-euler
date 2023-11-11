package main

import (
	"fmt"
	"slices"
)

func triangleFormula(n int) int {
	return (n * (n + 1)) / 2
}

func squareFormula(n int) int {
	return n * n
}

func pentagonalFormula(n int) int {
	return (n * (3*n - 1)) / 2
}

func hexagonalFormula(n int) int {
	return (n * (2*n - 1))
}

func heptagonalFormula(n int) int {
	return (n * (5*n - 3)) / 2
}

func octagonalFormula(n int) int {
	return n * (3*n - 2)
}

func polygonalFormula(polygonal, n int) int {
	if polygonal == 3 {
		return triangleFormula((n))
	} else if polygonal == 4 {
		return squareFormula(n)
	} else if polygonal == 5 {
		return pentagonalFormula(n)
	} else if polygonal == 6 {
		return hexagonalFormula(n)
	} else if polygonal == 7 {
		return heptagonalFormula(n)
	}

	return octagonalFormula(n)
}

func firstTwoDigits(n int) int {
	for n > 100 {
		n = n / 100
	}

	return n
}

func lastTwoDigits(n int) int {
	for n > 100 {
		n = n % 100
	}

	return n
}

func areCyclics(ints []int, minLength int) (bool, []int) {
	if len(ints) <= 1 {
		return false, []int{}
	}

	matcheds := []int{}

	for i := 0; i < 1; i++ {
		matcheds = []int{}

		number1 := ints[i]

		for {
			matched := false

			firstTwoDigits := firstTwoDigits(number1)

			if len(matcheds)+1 == len(ints) {
				lastTwoDigits := lastTwoDigits(matcheds[0])

				if firstTwoDigits == lastTwoDigits {
					matcheds = append(matcheds, number1)
				}
			} else {
				for _, number2 := range ints {
					if number1 == number2 || !has4Digits(number2) {
						continue
					}

					if slices.Contains(matcheds, number2) {
						continue
					}

					lastTwoDigits := lastTwoDigits(number2)

					if firstTwoDigits == lastTwoDigits {
						matcheds = append(matcheds, number1)
						number1 = number2
						matched = true
						break
					}
				}
			}

			if !matched {
				break
			}
		}

		if len(matcheds) >= minLength {
			break
		}
	}

	return len(matcheds) >= minLength, matcheds
}

func has4Digits(n int) bool {
	return n >= 1000 && n <= 9999
}

func hasMoreThan4Digits(n int) bool {
	return n > 9999
}

func checkCyclics(ns, polygonals, acc []int, results map[int]map[int]int) [][]int {
	if len(acc) == len(results) {
		cool, matches := areCyclics(acc, len(acc))
		if cool {
			fmt.Println("ns", ns)
			fmt.Println("Cyclics", acc)

			return [][]int{matches}
		}

		return [][]int{}
	}

	result := [][]int{}

	for idx, p := range polygonals {
		polygonalResults := results[p]

		for n, nResult := range polygonalResults {
			if slices.Contains(ns, n) {
				continue
			}

			if len(acc) > 0 {
				candidate := false

				nFirstTwoDigits := firstTwoDigits(nResult)
				nLastTwoDigits := lastTwoDigits(nResult)

				for _, inAcc := range acc {
					accFirstTwoDigits := firstTwoDigits(inAcc)
					accLastTwoDigits := lastTwoDigits(inAcc)

					if accLastTwoDigits == nFirstTwoDigits || accFirstTwoDigits == nLastTwoDigits {
						candidate = true
						break
					}
				}

				if !candidate {
					continue
				}
			}

			innerAccResult := checkCyclics(append(ns, n), polygonals[1:], append(acc, nResult), results)

			if len(innerAccResult) > 0 {
				result = append(result, innerAccResult[0])
			}
		}

		if len(acc) != idx+1 {
			break
		}
	}

	return result

}

func main() {
	polygonals := []int{3, 4, 5, 6, 7, 8}

	results := map[int]map[int]int{}

	for _, p := range polygonals {
		pResults := map[int]int{}

		for n := 1; ; n++ {
			pResult := polygonalFormula(p, n)

			if hasMoreThan4Digits(pResult) {
				break
			}

			if !has4Digits(pResult) {
				continue
			}

			pResults[n] = pResult

		}

		results[p] = pResults
	}

	allResults := checkCyclics([]int{}, polygonals, []int{}, results)

	sum := 0

	for _, n := range allResults[0] {
		sum += n
	}

	fmt.Println(allResults)
	fmt.Println(sum)

}
