package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func isValidRomanNumeral(numeral string) bool {
	uniqueLetters := [3]string{"D", "L", "V"}

	for _, letter := range uniqueLetters {
		if strings.Count(numeral, letter) > 1 {
			return false
		}
	}

	return true
}

func main() {
	shortenerMap := map[string]string{
		"DCCCC": "CM",
		"CCCCC": "D",
		"CCCC":  "CD",
		"LXXXX": "XC",
		"XXXXX": "L",
		"XXXX":  "XL",
		"VIIII": "IX",
		"IIIII": "V",
		"IIII":  "IV",
	}

	dat, err := os.ReadFile("./0089_roman.txt")

	check(err)

	rawContent := string(dat)

	savedNumbers := 0

	for _, rawNumber := range strings.Split(rawContent, "\n") {

		modifiedNumber := rawNumber

		for key, value := range shortenerMap {
			if strings.Contains(modifiedNumber, key) {
				result := strings.ReplaceAll(modifiedNumber, key, value)

				if isValidRomanNumeral(result) {
					modifiedNumber = result
				}
			}
		}

		if strings.Compare(rawNumber, modifiedNumber) != 0 {
			fmt.Print(rawNumber)
			fmt.Print(" ->> ")
			fmt.Println(modifiedNumber)

			savedNumbers += (len(rawNumber) - len(modifiedNumber))
		}
	}

	fmt.Println(savedNumbers)
}
