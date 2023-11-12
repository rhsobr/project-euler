package primes

import (
	"math/big"
)

func recursiveGreatestCommonFactor(number1 int, number2 int) int {
	if number2 == 0 {
		return number1
	}

	if number1 < number2 {
		return recursiveGreatestCommonFactor(number2, number1)
	}

	return recursiveGreatestCommonFactor(number2, number1%number2)
}

func AreRelativePrimes(number1 int, number2 int) bool {
	if number1%2 == 0 && number2%2 == 0 {
		return false
	}

	return recursiveGreatestCommonFactor(number1, number2) == 1
}

var primesMap map[int]bool = make(map[int]bool)

func IsPrime(number int) bool {
	result, ok := primesMap[number]

	if ok {
		return result
	}

	primesMap[number] = big.NewInt(int64(number)).ProbablyPrime(0)

	return primesMap[number]
}

var nextPrimesMap map[int]int = make(map[int]int)

func NextPrime(number int) int {
	result, ok := nextPrimesMap[number]

	if ok {
		return result
	}

	for i := number + 1; ; i++ {
		if IsPrime(i) {
			nextPrimesMap[number] = i
			return i
		}
	}
}
