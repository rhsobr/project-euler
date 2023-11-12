package main

import (
	"fmt"
	"project_euler/utils"

	"golang.org/x/exp/slices"
)

var factorialsSumOfDigitsMap map[int]int = make(map[int]int)

func getFactorialSumOfDigits(n int) (sum int) {
	mSum, ok := factorialsSumOfDigitsMap[n]

	if ok {
		return mSum
	}

	for _, digit := range utils.GetDigits(n) {
		sum += utils.GetFactorial(digit)
	}

	factorialsSumOfDigitsMap[n] = sum

	return sum
}

func getDigitFactorialChainsWithAcc(n int, acc []int) []int {
	sum := getFactorialSumOfDigits(n)

	if slices.Contains(acc, sum) {
		return acc
	}

	return getDigitFactorialChainsWithAcc(sum, append(acc, sum))
}

func getDigitFactorialChains(n int) []int {
	return getDigitFactorialChainsWithAcc(n, []int{n})
}

func main() {
	counter := 0

	for i := 1; i < 1_000_000; i++ {
		chains := getDigitFactorialChains(i)

		if len(chains) == 60 {
			counter++
		}
	}

	fmt.Println("Counter", counter)
}
