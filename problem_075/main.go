package main

import (
	"fmt"
	"math"
	"project_euler/utils/primes"
)

func getSides(m, n int) (a, b, c int) {
	if m < n {
		return getSides(n, m)
	}

	a = (m * m) - (n * n)
	b = (2 * m * n)
	c = (m * m) + (n * n)

	return a, b, c
}

func getPerimeter(m, n int) int {
	if m < n {
		return getPerimeter(n, m)
	}

	return (2 * m * m) + (2 * m * n)
}

var uniqueSquareTringleMap map[int]int = make(map[int]int)

func main() {
	cm := 1_500_000

	for m := 2; m <= int(math.Sqrt(float64(cm))); m++ {
		for n := 1; n < m; n++ {
			if (m+n)%2 != 1 {
				continue
			}

			if !primes.AreRelativePrimes(m, n) {
				continue
			}

			a, b, c := getSides(m, n)

			perimeter := a + b + c

			multipler := 1

			for {
				mult := multipler * perimeter

				if mult > cm {
					break
				}

				_, ok := uniqueSquareTringleMap[mult]

				if !ok {
					uniqueSquareTringleMap[mult] = 0
				}

				uniqueSquareTringleMap[mult]++

				multipler++
			}
		}
	}

	count := 0

	for ref, repetitions := range uniqueSquareTringleMap {
		if ref <= cm && repetitions == 1 {
			count++
		}
	}

	fmt.Println("Final count", count)

}
