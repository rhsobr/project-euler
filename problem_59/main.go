package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dat, _ := os.ReadFile("./0059_cipher.txt")

	rawContent := string(dat)

	numbers := strings.Split(rawContent, ",")

	f, _ := os.Create("temp.txt")

	defer f.Close()

	aASCII := 97
	zASCII := 123

	for i := aASCII; i < zASCII; i++ {
		for j := aASCII; j < zASCII; j++ {
			for h := aASCII; h < zASCII; h++ {
				if i == 101 && j == 120 && h == 112 {
					keys := []int{i, j, h}

					var data []int

					for pos, n := range numbers {

						var letterAscii int

						letterCryptedInteger, _ := strconv.Atoi(n)

						letterAscii = letterCryptedInteger ^ keys[pos%3]

						data = append(data, letterAscii)

					}

					var line string

					for _, r := range data {
						if 32 <= r && r <= 126 {
							line += string(r)
						} else {
							line = ""
							break
						}
					}

					sum := 0
					for _, x := range line {
						sum += int(x)
					}

					fmt.Println(sum)

					if len(line) > 0 {
						f.WriteString(string(i) + string(j) + string(h) + " - " + line + "\n")
					}
				}
			}
		}

	}
}
