package main

import (
	"fmt"
	"math/big"
	"time"
)

const expansions = 1000

func main() {
	start := time.Now()
	var (
		num         = big.NewInt(1)
		denom       = big.NewInt(2)
		bigTwo      = big.NewInt(2)
		tempDenom   = new(big.Int)
		bigNumTotal = 0
	)
	for i := 1; i < expansions; i++ {
		tempDenom = tempDenom.Mul(denom, bigTwo).Add(tempDenom, num)
		num.Set(denom)
		denom.Set(tempDenom)
		if len(tempDenom.Add(num, denom).String()) > len(denom.String()) {
			bigNumTotal++
		}
	}
	fmt.Printf("Took %.2f ms\n", time.Since(start).Seconds()*1000)
	fmt.Printf("Number of numerators bigger than denominators: %d\n", bigNumTotal)
}
