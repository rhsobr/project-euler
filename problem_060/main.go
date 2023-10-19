package main

import (
	"fmt"
	"project_euler/utils/primes"
)

func concatenateIntegers(int1 int, int2 int) int {
	shift := 10

	for shift <= int2 {
		shift *= shift
	}

	return (int1 * shift) + int2
}

func areConcatenatedPrimes(int1 int, int2 int) bool {

	integer1 := concatenateIntegers(int1, int2)

	if !primes.IsPrime(integer1) {
		return false
	}

	integer2 := concatenateIntegers(int2, int1)

	if !primes.IsPrime(integer2) {
		return false
	}

	return true
}

func main() {
	maxPrime := 9000
	minSum := 100000

	var primesCollection []int

	for i := 3; i <= maxPrime; i = primes.NextPrime(i) {
		primesCollection = append(primesCollection, i)
	}

	for i := 0; i < len(primesCollection); i++ {
		prime1 := primesCollection[i]
		if prime1 == 5 {
			continue
		}
		for j := len(primesCollection) - 1; j > (i + 4); j-- {
			prime5 := primesCollection[j]
			if prime5 == 5 {
				continue
			}
			if prime1+prime5 >= minSum {
				break
			}

			if !areConcatenatedPrimes(prime1, prime5) {
				continue
			}

			for h := i + 1; h < j; h++ {
				prime3 := primesCollection[h]

				if prime1+prime5+prime3 >= minSum {
					break
				}

				if !areConcatenatedPrimes(prime3, prime5) || !areConcatenatedPrimes(prime1, prime3) {
					continue
				}

				for z := h + 1; z < j; z++ {
					prime4 := primesCollection[z]

					if prime1+prime5+prime3+prime4 >= minSum {
						break
					}

					if !areConcatenatedPrimes(prime3, prime4) || !areConcatenatedPrimes(prime5, prime4) || !areConcatenatedPrimes(prime1, prime4) {
						continue
					}

					for o := z + 1; o < j; o++ {
						prime2 := primesCollection[o]
						sum := prime1 + prime2 + prime3 + prime4 + prime5

						if sum >= minSum {
							break
						}

						if !areConcatenatedPrimes(prime1, prime2) || !areConcatenatedPrimes(prime2, prime5) || !areConcatenatedPrimes(prime3, prime2) || !areConcatenatedPrimes(prime4, prime2) {
							continue
						}

						minSum = sum
						primesArr := []int{prime1, prime2, prime3, prime4, prime5}
						fmt.Println(primesArr, sum)

					}
				}
			}

		}
	}

	fmt.Println(minSum)
}
