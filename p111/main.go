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
const numDigits = 10

// This function uses a segmented sieve to find repeated digits in primes
// with a certain number of digits
func repeat_segmented_sieve(lowerLimit, limit int64, digits map[int64]*RepeatDigits) {
	var (
		primes    []int64
		multiples []int64
	)
	sqrt := int64(math.Sqrt(float64(limit)))
	segmentSize := max(sqrt, L1D_CACHE_SIZE)

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
			if sieve[n-low] && n > lowerLimit { // n is a prime
				num, numRepeated := findRepitition(n)
				// Ignore all numbers that don't have enough repititive numbers
				if numRepeated < numDigits-2 {
					continue
				}
				if digits[num].NumRepeated < numRepeated {
					digits[num].NumRepeated = numRepeated
					digits[num].Sum = float64(n)
					//digits[num].Primes = []int64{n}
				} else if digits[num].NumRepeated == numRepeated {
					digits[num].Sum += float64(n)
					//digits[num].Primes = append(digits[num].Primes, i)
				}
			}
		}
	}
}

type RepeatDigits struct {
	NumRepeated int
	Primes      []int64
	Sum         float64
}

var repititionMap = map[int64]int{
	0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0, 9: 0,
}

func findRepitition(x int64) (int64, int) {
	var mod int64
	var max int
	var maxIndex int64
	for x > 0 {
		mod = x % 10
		repititionMap[mod]++
		x /= 10
		if repititionMap[mod] > max {
			max = repititionMap[mod]
			maxIndex = mod
		}
	}
	for k := range repititionMap {
		repititionMap[k] = 0
	}
	return maxIndex, max
}

func main() {
	start := int64(1)
	for i := 0; i < numDigits-1; i++ {
		start *= 10
	}
	end := start * 10
	start++
	var digits = make(map[int64]*RepeatDigits, 0)
	for i := int64(0); i < 10; i++ {
		digits[i] = &RepeatDigits{}
	}

	repeat_segmented_sieve(start, end, digits)

	total := float64(0)
	for _, repeatDigit := range digits {
		// fmt.Printf("Key: %d\n", k)
		// print("Primes: ")
		// for _, prime := range repeatDigit.Primes {
		// 	fmt.Printf("%d, ", prime)
		// }
		total += repeatDigit.Sum
		// println()
	}
	fmt.Printf("%.f\n", total)
	//println(otherTotal)

}

// func main() {
// 	digits := make(map[int64]*RepeatDigits, 0)
// 	for i := int64(0); i < 10; i++ {
// 		digits[i] = &RepeatDigits{}
// 	}
// 	prime := 0
// 	for i := int64(10000001); i < 100000000; i += 2 {
// 		if millerRabin(i) {
// 			prime++
// 		}
// 	}
// 	fmt.Printf("There are %d priems\n", prime)
// }
