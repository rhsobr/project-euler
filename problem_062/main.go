package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func containsAllDigits(numberStr1 string, numberStr2 string) bool {
	if len(numberStr1) != len(numberStr2) {
		return false
	}

	for _, digit := range strings.Split(numberStr1, "") {
		if strings.Count(numberStr1, digit) != strings.Count(numberStr2, digit) {
			return false
		}
	}

	return true
}

func doCheck(str1 string, str2 string) bool {
	return str1 != str2 && containsAllDigits(str1, str2)
}

func main() {
	cubeMap := map[int]string{}

	for i := 9999; i >= 345; i-- {
		cube := math.Pow(float64(i), 3)
		str := strconv.Itoa(int(cube))
		cubeMap[i] = str
	}

	bag := []string{}

	for _, str1 := range cubeMap {
		for _, str2 := range cubeMap {
			if !doCheck(str1, str2) {
				continue
			}

			for _, str3 := range cubeMap {
				if !doCheck(str3, str2) || !doCheck(str3, str1) {
					continue
				}

				for _, str4 := range cubeMap {
					if !doCheck(str4, str3) || !doCheck(str4, str2) || !doCheck(str4, str1) {
						continue
					}

					for _, str5 := range cubeMap {
						if !doCheck(str5, str4) || !doCheck(str5, str3) || !doCheck(str5, str2) || !doCheck(str5, str1) {
							continue
						}

						bag = append(bag, str1)
					}
				}
			}
		}
	}

	min := 0

	for _, x := range bag {
		xs, _ := strconv.Atoi(x)

		if min == 0 || min >= xs {
			min = xs
		}
	}

	fmt.Println(min)
}
