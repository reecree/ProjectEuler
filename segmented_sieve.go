package main

import (
	"fmt"
	"math"
)

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

// Set your CPU's L1 data cache size (in bytes) here
const L1D_CACHE_SIZE = int64(32768)

// This function uses a segmented sieve to find repeated digits in primes
// with a certain number of digits
func segmented_sieve(limit int64) {
	var (
		count     int64
		primes    []int64
		multiples []int64
	)
	sqrt := int64(math.Sqrt(float64(limit)))
	segmentSize := max(sqrt, L1D_CACHE_SIZE)
	if limit < 2 {
		count = 0
	} else {
		count = 1
	}

	sieve := make([]bool, segmentSize)
	isPrime := make([]bool, sqrt+1)
	for i := 0; i < len(isPrime); i++ {
		isPrime[i] = true
	}

	var i, n, s = int64(3), int64(3), int64(3)

	for low := int64(0); low <= limit; low += segmentSize {
		for i := 0; i < len(sieve); i++ {
			sieve[i] = true
		}

		// current segment = [low, high]
		high := low + segmentSize - 1
		high = min(high, limit)

		// generate sieving primes using simple sieve of Eratosthenes
		for ; i*i <= high; i += 2 {
			if isPrime[i] {
				for j := i * i; j <= sqrt; j += i {
					isPrime[j] = false
				}
			}
		}

		// initialize sieving primes for segmented sieve
		for ; s*s <= high; s += 2 {
			if isPrime[s] {
				primes = append(primes, s)
				multiples = append(multiples, s*s-low)
			}
		}

		// sieve the current segment
		for x := 0; x < len(primes); x++ {
			j := multiples[x]
			for k := primes[x] * 2; j < segmentSize; j += k {
				sieve[j] = false
			}
			multiples[x] = j - segmentSize
		}

		for ; n <= high; n += 2 {
			if sieve[n-low] { // n is a prime
				count++
			}
		}
	}

	fmt.Printf("%d primes found\n", count)
}
