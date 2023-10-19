package primes

import "math/big"

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

func IsPrime(number int) bool {
	return big.NewInt(int64(number)).ProbablyPrime(0)
}

var primesMap map[int]int = make(map[int]int)

func NextPrime(number int) int {
	result, ok := primesMap[number]

	if ok {
		return result
	}

	for i := number + 1; ; i++ {
		if IsPrime(i) {
			primesMap[number] = i
			return i
		}
	}
}
