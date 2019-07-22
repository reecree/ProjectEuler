package main

// import (
// 	"fmt"
// 	"strconv"
// 	"time"
// )

// const (
// 	LyBelowNum = 10000
// 	MaxTry     = 50
// )

// func IsPalindromic(x float64) (float64, bool) {
// 	var temp byte
// 	xstr := []byte(fmt.Sprintf("%.0f", x))
// 	flipStr := xstr
// 	isPalindrome := true
// 	for i := 0; i < len(xstr)/2; i++ {
// 		if xstr[i] != xstr[len(xstr)-i-1] {
// 			isPalindrome = false
// 			temp = flipStr[i]
// 			flipStr[i] = flipStr[len(flipStr)-i-1]
// 			flipStr[len(flipStr)-i-1] = temp
// 		}
// 	}

// 	flipX, _ := strconv.ParseFloat(string(flipStr), 64)
// 	return flipX, isPalindrome
// }

// func IsLychrel(x float64) bool {
// 	var flipX float64
// 	var result bool
// 	flipX, _ = IsPalindromic(x)
// 	x += flipX
// 	for i := 1; i < MaxTry; i++ {
// 		if flipX, result = IsPalindromic(x); result {
// 			return false
// 		}
// 		x += flipX
// 	}
// 	return true
// }

// func main() {
// 	numLy := 0
// 	var i float64
// 	start := time.Now()
// 	for i = 1; i < LyBelowNum; i++ {
// 		if IsLychrel(i) {
// 			numLy++
// 		}
// 	}
// 	fmt.Printf("%.5f\n", time.Since(start).Seconds())
// 	println(numLy)
// }
