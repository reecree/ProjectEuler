package main

import (
	"fmt"
	"math/big"
	"time"
)

const (
	Max = uint64(100)
)

var (
	BigZero   = big.NewInt(0)
	BigTen    = big.NewInt(10)
	remainder = new(big.Int)
	hold      = new(big.Int)
)

func ComputeDigitalSumRem(x *big.Int) *big.Int {
	sum := new(big.Int)
	for x.Cmp(BigZero) > 0 {
		sum.Add(sum, remainder.Rem(x, BigTen))
		x.Div(x, BigTen)
	}
	return sum
}

func ComputeDigitalDivMod(x *big.Int) *big.Int {
	sum := new(big.Int)
	for x.Cmp(BigZero) > 0 {
		hold.DivMod(x, BigTen, remainder)
		sum.Add(sum, remainder)
		x.Div(x, BigTen)
	}
	return sum
}

func LargestSum(
	maxA, maxB uint64,
	sumCompute func(*big.Int) *big.Int,
) *big.Int {
	var a, b, sum, result = new(big.Int), new(big.Int), new(big.Int), new(big.Int)
	largest := big.NewInt(0)
	for i := uint64(2); i < maxA; i++ {
		a.SetUint64(i)
		for j := uint64(2); j < maxB; j++ {
			b.SetUint64(j)
			result.Exp(a, b, nil)
			sum = sumCompute(result)
			if sum.Cmp(largest) > 0 {
				largest.Set(sum)
			}
		}
	}
	return largest
}

func main() {

	startRem := time.Now()
	largest := LargestSum(Max, Max, ComputeDigitalSumRem)
	fmt.Printf("Largest Sum %s w/ Rem took %.4f sec\n",
		largest.String(),
		time.Since(startRem).Seconds(),
	)

	startDivMod := time.Now()
	largest = LargestSum(Max, Max, ComputeDigitalDivMod)
	fmt.Printf("Largest Sum %s w/ Div Mod took %.4f sec\n",
		largest.String(),
		time.Since(startDivMod).Seconds(),
	)
}
