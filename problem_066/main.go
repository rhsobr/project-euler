package main

import (
	"fmt"
	"math/big"
	"project_euler/utils"
)

func main() {
	largestX := big.NewInt(1)
	matchD := 0

	for D := 3; D <= 1000; D++ {
		if utils.IsPerfectSquare(D) {
			continue
		}

		aks := []*big.Int{}
		pks := []*big.Int{}
		qks := []*big.Int{}

		solved := false

		for k := 0; k <= 1000; k++ {
			akTemp, pr, qr := utils.GetContinueFractionPosition(D, k)

			ak := big.NewInt(int64(akTemp))

			fmt.Println(akTemp, pr, qr)

			aks = append(aks, ak)

			// https://proofwiki.org/wiki/Continued_Fraction_Expansion_of_Irrational_Square_Root/Examples/61/Convergents
			if k == 0 {
				pks = append(pks, aks[k])
				qks = append(qks, big.NewInt(1))
			} else if k == 1 {
				var ref big.Int
				pks = append(pks, ref.Add(ref.Mul(aks[k-1], aks[k]), big.NewInt(1)))
				qks = append(qks, aks[k])
			} else {
				var ref big.Int
				var ref2 big.Int

				pks = append(pks, ref.Add(ref.Mul(aks[k], pks[k-1]), pks[k-2]))
				qks = append(qks, ref2.Add(ref2.Mul(aks[k], qks[k-1]), qks[k-2]))
			}

			x := pks[len(pks)-1]
			y := qks[len(qks)-1]

			if D%2 == 0 && big.NewInt(0).Mod(x, big.NewInt(2)).Cmp(big.NewInt(0)) == 0 {
				continue
			}

			var powX big.Int
			powX.Exp(x, big.NewInt(2), nil)

			var powY big.Int
			powY.Exp(y, big.NewInt(2), nil)
			powY.Mul(&powY, big.NewInt(int64(D)))

			powX.Sub(&powX, &powY)

			if powX.Cmp(big.NewInt(1)) == 0 {
				if largestX.Cmp(x) == -1 {
					fmt.Printf("%v %v %v\n", x, D, y)
					matchD = D
					largestX = x
				}
				solved = true
				break
			}
			//}

		}

		if !solved {
			fmt.Println("Not solved", D)
		}
	}

	fmt.Printf("Largest X %v\n", largestX)
	fmt.Println("Match D", matchD)

}
