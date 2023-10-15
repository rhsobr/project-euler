package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func isPerfectSquare(number int) bool {
	sqtr := math.Sqrt(float64(number))
	return int(sqtr)*int(sqtr) == number
}

func convertToArrayInt(arr []string) []int {
	var result []int
	for _, n := range arr {
		number, _ := strconv.Atoi(n)
		result = append(result, number)
	}
	return result
}

func sumArray(arr []int) int {
	nSum := 0
	for _, n := range arr {
		nSum += n
	}
	return nSum
}

func initRecursiveCombinations(digits []string, sum int, limit int) []int {
	var acc []int

	if sum > limit {
		return acc
	}

	lenDigits := len(digits)

	if lenDigits == 0 && sum == limit {
		return append(acc, sum)
	}

	for i := 1; i <= lenDigits; i++ {
		newNumberInt, _ := strconv.Atoi(strings.Join(digits[:i], ""))

		tail := digits[i:]

		remaining, _ := strconv.Atoi(strings.Join(tail, ""))

		if len(tail) == 0 || remaining > 0 {
			tempAcc := initRecursiveCombinations(tail, sum+newNumberInt, limit)

			if len(tempAcc) > 0 {
				return tempAcc
			}

			continue
		}

		tempAcc := initRecursiveCombinations(tail[0:1], sum+newNumberInt, limit)

		if len(tempAcc) > 0 {
			return tempAcc
		}

	}

	return acc
}

func isSNumber(number int, sqtr int) bool {
	digits := strings.Split(strconv.Itoa(number), "")

	acc := initRecursiveCombinations(digits, 0, sqtr)

	for _, sum := range acc {
		if sum == int(sqtr) {
			return true
		}
	}

	return false
}

func sumOfSNumbers(N int) int {
	nSum := 0

	for i := 2; ; i++ {
		pow := int(math.Pow(float64(i), 2))

		if pow >= N+1 {
			break
		}

		if isSNumber(pow, i) {
			fmt.Print(i)
			fmt.Print(" ---> ")
			fmt.Println(pow)
			nSum += pow
		}
	}
	return nSum
}

func main() {
	fmt.Println(sumOfSNumbers(int(math.Pow(10, 10))))
}
