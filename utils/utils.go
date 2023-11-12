package utils

import (
	"math"
	"math/big"
	"project_euler/utils/primes"
)

var allFactorsMap map[int]map[int]int = make(map[int]map[int]int)

func GetAllFactors(n int) map[int]int {
	cached, ok := allFactorsMap[n]

	if ok {
		return cached
	}

	factorsMap := map[int]int{}

	if n == 1 {
		return factorsMap
	}

	if primes.IsPrime(n) {
		factorsMap[n] = 1
		allFactorsMap[n] = factorsMap
		return factorsMap
	}

	currentPrime := 2

	for {
		if n%currentPrime > 0 {
			currentPrime = primes.NextPrime(currentPrime)
			continue
		}

		_, ok := factorsMap[currentPrime]

		if ok {
			factorsMap[currentPrime] += 1
		} else {
			factorsMap[currentPrime] = 1
		}

		for factor, count := range GetAllFactors(n / currentPrime) {
			_, ok := factorsMap[factor]

			if !ok {
				factorsMap[factor] = count
				continue
			}

			factorsMap[factor] += count
		}

		break
	}

	allFactorsMap[n] = factorsMap

	return factorsMap
}

func CountNumberOfRelativePrimes(number int) int {
	if primes.IsPrime(number) {
		return number - 1
	}

	factors := GetAllFactors(number)

	sum := 1

	for factor, count := range factors {
		fFactor := float64(factor)
		sum *= int(math.Pow(fFactor, float64(count)) - math.Pow(fFactor, float64(count-1)))
	}

	return sum
}

func getAnTestRec(initialNumber int, a float64, x float64, position int, currPosition int) []float64 {
	if position == currPosition {
		return []float64{a, x, (x - a)}
	}

	root := math.Sqrt(float64(initialNumber))

	a0 := math.Floor(root)

	if a0 > 1 && position > 1 && int(a)%int(a0) == 0 {
		position1 := GetContinueFractionPosition2(initialNumber, 1)
		return getAnTestRec(initialNumber, position1[0], position1[1], position, currPosition+1)
	}

	xNumerator := big.NewFloat(root + a)

	xDenominator := new(big.Float)
	xDenominator.Mul(big.NewFloat(x-a), big.NewFloat(root+a))

	xPosition := xNumerator.Quo(xNumerator, xDenominator)

	xPositionFloat64, _ := xPosition.Float64()

	aPosition := math.Floor(xPositionFloat64)

	return getAnTestRec(initialNumber, aPosition, xPositionFloat64, position, currPosition+1)

}

func GetContinueFractionPosition2(n int, position int) []float64 {
	root := math.Sqrt(float64(n))

	a0 := math.Floor(root)

	return getAnTestRec(n, a0, root, position, 0)
}

// https://proofwiki.org/wiki/Continued_Fraction_Expansion_of_Irrational_Square_Root/Examples/61
func GetContinueFractionPosition(n int, position int) (int, int, int) {

	if position == 0 {
		return int(math.Floor((math.Sqrt(float64(n))))), 0, 1
	}
	var previousA, previousPR, previousQR int

	if position > 0 {
		previousA, previousPR, previousQR = GetContinueFractionPosition(n, position-1)
	} else {
		previousA = 0
		previousPR = 0
		previousQR = 1
	}

	pr := float64(previousA*previousQR - previousPR)
	qr := (float64(n) - math.Pow(float64(pr), 2)) / float64(previousQR)

	ar := int(math.Floor((math.Sqrt(float64(n)) + pr) / qr))

	return ar, int(pr), int(qr)
}

func IsPerfectSquare(n int) bool {
	sqrt := math.Sqrt(float64(n))
	return sqrt-math.Floor(sqrt) == 0
}

var factorialsMap map[int]int = make(map[int]int)

func GetFactorial(n int) int {
	memo, ok := factorialsMap[n]

	if ok {
		return memo
	}

	if n > 0 {
		factorialsMap[n] = n * GetFactorial(n-1)
		return factorialsMap[n]
	}

	return 1
}

func getDigitsWithAcc(n int, acc []int) []int {
	if n < 10 {
		return append([]int{n}, acc...)
	}

	return getDigitsWithAcc(n/10, append([]int{n % 10}, acc...))
}

func GetDigits(n int) []int {
	return getDigitsWithAcc(n, []int{})
}
