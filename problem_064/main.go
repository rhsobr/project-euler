package main

import (
	"fmt"
	"math"
	"project_euler/utils"
)

func main() {
	greaterOddPeriod := 0

	for i := 2; i <= 10000; i++ {
		if utils.IsPerfectSquare(i) {
			continue
		}

		root := math.Sqrt(float64(i))
		a0 := int(math.Floor(root))

		continuedFraction := []int{}

		for pos := 1; ; pos++ {
			aPosition, _, _ := utils.GetContinueFractionPosition(i, pos)

			continuedFraction = append(continuedFraction, aPosition)

			if aPosition > a0 && aPosition%a0 == 0 {
				if pos%2 == 1 {
					greaterOddPeriod++

				}
				break
			}
		}
	}

	fmt.Println(greaterOddPeriod)

}
