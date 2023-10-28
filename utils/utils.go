package utils

import (
	"math"
	"math/big"
)

func getAnTestRec(initialNumber int, a float64, x float64, position int, currPosition int) []float64 {
	if position == currPosition {
		return []float64{a, x, (x - a)}
	}

	root := math.Sqrt(float64(initialNumber))

	a0 := math.Floor(root)

	if a0 > 1 && position > 1 && int(a)%int(a0) == 0 {
		position1 := GetContinueFractionPosition2(initialNumber, 1)
		return getAnTestRec(initialNumber, position1[0], position1[1], position, currPosition+1)
	}

	xNumerator := big.NewFloat(root + a)

	xDenominator := new(big.Float)
	xDenominator.Mul(big.NewFloat(x-a), big.NewFloat(root+a))

	xPosition := xNumerator.Quo(xNumerator, xDenominator)

	xPositionFloat64, _ := xPosition.Float64()

	aPosition := math.Floor(xPositionFloat64)

	return getAnTestRec(initialNumber, aPosition, xPositionFloat64, position, currPosition+1)

}

func GetContinueFractionPosition2(n int, position int) []float64 {
	root := math.Sqrt(float64(n))

	a0 := math.Floor(root)

	return getAnTestRec(n, a0, root, position, 0)
}

// https://proofwiki.org/wiki/Continued_Fraction_Expansion_of_Irrational_Square_Root/Examples/61
func GetContinueFractionPosition(n int, position int) (int, int, int) {

	if position == 0 {
		return int(math.Floor((math.Sqrt(float64(n))))), 0, 1
	}
	var previousA, previousPR, previousQR int

	if position > 0 {
		previousA, previousPR, previousQR = GetContinueFractionPosition(n, position-1)
	} else {
		previousA = 0
		previousPR = 0
		previousQR = 1
	}

	pr := float64(previousA*previousQR - previousPR)
	qr := (float64(n) - math.Pow(float64(pr), 2)) / float64(previousQR)

	ar := int(math.Floor((math.Sqrt(float64(n)) + pr) / qr))

	return ar, int(pr), int(qr)
}

func IsPerfectSquare(n int) bool {
	sqrt := math.Sqrt(float64(n))
	return sqrt-math.Floor(sqrt) == 0
}
