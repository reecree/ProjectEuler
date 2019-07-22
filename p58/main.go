package main

import (
	"fmt"
	"time"
)

const (
	numCorners = 4
)

// Based on https://en.wikipedia.org/wiki/Miller%E2%80%93Rabin_primality_test
func millerRabin(n uint64) bool {
	firstPrimes := []uint64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79}
	witnesses := []uint64{2, 7, 61}
	if n < 2 {
		return false
	}
	for _, p := range firstPrimes {
		if n%p == 0 {
			return n == p
		}
	}

	s := uint64(0)
	d := uint64(n - 1)
	for d%2 == 0 && d > 0 {
		d /= 2
		s++
	}

	var (
		power uint64
		found bool
		a     uint64
	)

	for _, b := range witnesses {
		a = b % n
		if a == 0 {
			return true
		}
		power = intPowMod(a, d, n)
		if power == 1 || power == n-1 {
			continue
		}
		found = false
		for r := uint64(1); r < s; r++ {
			power *= power
			power %= n
			if power == 1 {
				return false
			}
			if power == n-1 {
				found = true
				break
			}
		}
		if found {
			continue
		}
		return false
	}
	return true
}

func intPowMod(base, exp, mod uint64) uint64 {
	result := uint64(1)
	for exp >= 2 {
		if exp%2 != 0 {
			result *= base
			if result >= mod {
				result %= mod
			}
		}
		base *= base
		if base >= mod {
			base %= mod
		}
		exp /= 2
	}
	result *= base
	result %= mod
	return result
}

func main() {
	diagnols := float32(5)
	primes := float32(3)
	i := uint64(9)
	multiple := uint64(4)
	sideLength := 3
	start := time.Now()
	for primes/diagnols > .1 {
		for j := 0; j < numCorners; j++ {
			i += multiple
			diagnols++
			// j = 3 are all sqrs so they don't need to be checked
			if j != numCorners-1 && millerRabin(i) {
				primes++
			}
		}
		sideLength += 2
		multiple += 2
	}
	fmt.Printf("Took %.3f ms\n", time.Since(start).Seconds()*1000)
	fmt.Printf("SideLength when ratio below 10%% is %d\n", sideLength)
}
