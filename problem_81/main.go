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

func main() {
	dat, _ := os.ReadFile("./0081_matrix.txt")

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

	matrixLength := len(matrix) - 1

	for i := matrixLength; i >= 0; i-- {
		for j := matrixLength; j >= 0; j-- {
			if i < matrixLength && j < matrixLength {
				matrix[i][j] += int(math.Min(float64(matrix[i+1][j]), float64(matrix[i][j+1])))
			} else if i < matrixLength {
				matrix[i][j] += matrix[i+1][j]
			} else if j < matrixLength {
				matrix[i][j] += matrix[i][j+1]
			}
		}
	}

	fmt.Println(matrix[0][0])
}
