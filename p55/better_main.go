package main

import (
	"fmt"
	"strconv"
	"time"
)

const (
	LyBelowNum = 10000
	MaxTry     = 50
)

var (
	lychrels          = make(map[int]struct{}, 0)
	nonlychrels       = make(map[int]struct{}, 0)
	attemptedBelowMax = make([]int, 100)
)

func IsPalindromic(x float64) (float64, bool) {
	var temp byte
	xstr := []byte(fmt.Sprintf("%.0f", x))
	flipStr := xstr
	isPalindrome := true
	for i := 0; i < len(xstr)/2; i++ {
		if xstr[i] != xstr[len(xstr)-i-1] {
			isPalindrome = false
			temp = flipStr[i]
			flipStr[i] = flipStr[len(flipStr)-i-1]
			flipStr[len(flipStr)-i-1] = temp
		}
	}

	flipX, _ := strconv.ParseFloat(string(flipStr), 64)
	return flipX, isPalindrome
}

func IsLychrel(x float64) bool {
	_, exists := lychrels[int(x)]
	if exists {
		return true
	}
	var (
		result       bool
		flipX        float64
		numAttempted int
	)
	if flipX, result = IsPalindromic(x); !result {
		attemptedBelowMax[numAttempted] = int(flipX)
		numAttempted++
	}

	x += flipX
	for i := 1; i < MaxTry; i++ {
		// Need to check inside loop otherwise will misidentify
		// palindromic lychrels
		_, exists = nonlychrels[int(x)]
		if exists {
			return false
		}
		if x < LyBelowNum {
			attemptedBelowMax[numAttempted] = int(x)
			numAttempted++
		}
		if flipX, result = IsPalindromic(x); result {
			for i := 0; i < numAttempted; i++ {
				nonlychrels[attemptedBelowMax[i]] = struct{}{}
			}
			return false
		}
		if flipX < LyBelowNum {
			attemptedBelowMax[numAttempted] = int(flipX)
			numAttempted++
		}
		x += flipX
	}
	for i := 0; i < numAttempted; i++ {
		lychrels[attemptedBelowMax[i]] = struct{}{}
	}

	return true
}

func main() {
	start := time.Now()
	numLy := 0
	var i float64
	for i = 1; i < LyBelowNum; i++ {
		if IsLychrel(i) {
			numLy++
		}
	}
	fmt.Printf("Solution took %.5f seconds\n", time.Since(start).Seconds())
	fmt.Printf("%d lychrels below %d\n", numLy, LyBelowNum)
}
