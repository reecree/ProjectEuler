package main

import (
	"fmt"
	"math"
	"strconv"
)

// PrimeLimit < 100 million should be enough to generate 5 primes
var PrimeLimit = 100000000

// CheckPrimeLimit Needs to be < 10,000, due to our prime limit being
// 100 million
var CheckPrimeLimit = 10000
var NumberOfComboPrimes = 5

// sieveOfEratosthenes returns the entire sieve. As well as a list of only
// primes
// we should only be dealing with primes under 1 to 100 million. So
// sieve of erathonos should suffice
func sieveOfEratosthenes(limit int, primeMaxLimit int) ([]bool, []int) {
	if limit < 2 {
		return nil, []int{}
	}
	sieve := make([]bool, limit)
	// Could probably improve perforamce by using the prime number theorem
	// to predict approximately how many primes we'll have, but, as our
	// number of primes should be relatively small (order of 10,000s), not
	// doing for now
	primes := []int{}
	limitSqrt := math.Sqrt(float64(limit))
	for i := 2; i < int(limitSqrt); i++ {
		if !sieve[i] {
			if i < primeMaxLimit {
				primes = append(primes, i)
			}
			j := i * i
			for j < limit {
				sieve[j] = true
				j += i
			}
		}
	}
	for i := int(limitSqrt); i < limit; i++ {
		if !sieve[i] && i < primeMaxLimit {
			primes = append(primes, i)
		}
	}
	return sieve, primes
}

func checkPrimalityOfCombo(p1, p2 int, sieve []bool) bool {
	// Should be impossible to error as we started with two ints, this
	// would error if we overflowed the int, but that won't happen with
	// the size of primes we're dealing with
	comboP1, _ := strconv.Atoi(strconv.Itoa(p1) + strconv.Itoa(p2))
	if sieve[comboP1] {
		return false
	}
	comboP2, _ := strconv.Atoi(strconv.Itoa(p2) + strconv.Itoa(p1))
	return !sieve[comboP2]
}

func main() {

	sieve, primes := sieveOfEratosthenes(PrimeLimit, CheckPrimeLimit)
	fmt.Println("Finished generating sieve")

	// We should only need to check the first 100 or so primes as 'first level'
	// primes in order to find the smallest sum. That is pure conjecture
	// though and could very much be wrong. BUT if we do find any with
	// first level in the first 100, we can be pretty confident that the
	// sum will be the lowest
	for i := 1; i < 100; i++ {
		checkSet([]int{primes[i]}, primes, i+1, sieve)
	}
}

func checkSet(goodPrimes, allPrimes []int, startIndex int, sieve []bool) {
	// Should always check the biggest number first as that should prune the
	// tree quicker, so iterate in reverse order
	for i := startIndex; i < len(allPrimes); i++ {
		sum := allPrimes[i]
		allMatch := true
		for j := len(goodPrimes) - 1; j >= 0; j-- {
			if !checkPrimalityOfCombo(goodPrimes[j], allPrimes[i], sieve) {
				allMatch = false
				break
			}
			sum += goodPrimes[j]
		}
		if allMatch {
			if len(goodPrimes)+1 >= NumberOfComboPrimes {
				fmt.Printf(
					"Found Set of 5 combo primes: %+v\nSum: %d\n\n",
					append(goodPrimes, allPrimes[i]), sum)
			} else {
				checkSet(append(goodPrimes, allPrimes[i]), allPrimes, i+1, sieve)
			}
		}
	}
}
