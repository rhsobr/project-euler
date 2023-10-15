package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func sumArray(arr []int) int {
	nSum := 0

	for _, n := range arr {
		nSum += n
	}

	return nSum
}

func run(rowIdx int, columnIdx int, matrix [][]int, sum int, limit int) int {

	matrixLen := len(matrix)

	if rowIdx >= matrixLen-1 && columnIdx >= matrixLen-1 {
		return sum
	}

	if sum >= limit {
		return limit
	}

	if rowIdx >= matrixLen-1 {
		return run(rowIdx, columnIdx+1, matrix, sum+matrix[rowIdx][columnIdx+1], limit)
	}

	if columnIdx >= matrixLen-1 {
		return run(rowIdx+1, columnIdx, matrix, sum+matrix[rowIdx+1][columnIdx], limit)
	}

	sum1 := run(rowIdx+1, columnIdx, matrix, sum+matrix[rowIdx+1][columnIdx], limit)

	sum2 := run(rowIdx, columnIdx+1, matrix, sum+matrix[rowIdx][columnIdx+1], int(math.Min(float64(sum1), float64(limit))))

	if sum1 <= sum2 {
		return sum1
	}

	return sum2
}

func main() {
	dat, _ := os.ReadFile("./0081_sample3.txt")

	rawContent := string(dat)

	lines := strings.Split(rawContent, "\n")

	lenLines := len(lines)

	matrix := make([][]int, lenLines)

	for rowIdx, line := range lines {
		for _, number := range strings.Split(line, ",") {
			n, _ := strconv.Atoi(number)
			matrix[rowIdx] = append(matrix[rowIdx], n)
		}
	}

	rowIdx := 0
	columnIdx := 0
	var path []int

	for {
		path = append(path, matrix[rowIdx][columnIdx])

		if rowIdx >= lenLines-1 && columnIdx >= lenLines-1 {
			break
		} else if rowIdx >= lenLines-1 {
			columnIdx += 1
			continue
		} else if columnIdx >= lenLines-1 {
			rowIdx += 1
			continue
		}

		left := matrix[rowIdx][columnIdx+1]
		bottom := matrix[rowIdx+1][columnIdx]

		if left > bottom {
			rowIdx += 1
		} else {
			columnIdx += 1
		}
	}

	optimisticSum := sumArray(path)

	fmt.Println(path)

	sum := run(0, 0, matrix, matrix[0][0], optimisticSum)

	fmt.Println(sum)
}
