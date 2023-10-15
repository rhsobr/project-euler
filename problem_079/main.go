package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func exchangePositions(arr []int, pos1 int, pos2 int) []int {
	ref := arr[pos2]
	arr[pos2] = arr[pos1]
	arr[pos1] = ref
	return arr
}

func main() {
	data, _ := os.ReadFile("./0079_keylog.txt")

	numbers := strings.Split(string(data), "\n")

	var result []int

	for _, number := range numbers {
		integer, _ := strconv.Atoi(number)

		digit1 := (integer%1000 - (integer % 100)) / 100
		digit2 := (integer%100 - (integer % 10)) / 10
		digit3 := integer % 10

		for _, digits := range [][]int{{digit1, digit2}, {digit2, digit3}} {
			index1 := slices.Index(result, digits[0])
			index2 := slices.Index(result, digits[1])

			if index1 < 0 && index2 < 0 {
				result = append(result, digits[0])
				result = append(result, digits[1])
			} else if index1 < 0 && index2 > 0 {
				result = slices.Insert(result, index2, digits[0])
			} else if index1 > 0 && index2 < 0 {
				result = slices.Insert(result, index1+1, digits[1])
			} else if index1 > index2 {
				result = exchangePositions(result, index1, index2)
			}
		}
	}

	fmt.Println(result)
}
