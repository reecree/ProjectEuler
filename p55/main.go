package main

// import (
// 	"math/big"
// )

// type Lychrel struct {
// 	x *big.Int
// }

// const (
// 	LyBelowNum = 10000
// 	MaxTry     = 50
// )

// var (
// 	// BigInts necessary for temp storage when computing
// 	bigTen        = big.NewInt(10)
// 	bigHundred    = big.NewInt(100)
// 	hold          = big.NewInt(0)
// 	checkBegin    = big.NewInt(0)
// 	flipBegin     = big.NewInt(0)
// 	checkEnd      = big.NewInt(0)
// 	flipEnd       = big.NewInt(0)
// 	divisor       = big.NewInt(0)
// 	divisorY      = big.NewInt(0)
// 	divisorYUp    = big.NewInt(0)
// 	tempX         = big.NewInt(0)
// 	lyX           = big.NewInt(0)
// 	flipX         = big.NewInt(0)
// 	mod           = big.NewInt(0)
// 	bigLyBelowNum = big.NewInt(LyBelowNum)

// 	lychrels = make([]bool, LyBelowNum)
// 	result   = false

// 	maxLength = 0
// )

// func IsPalindromicString(x *big.Int) (*big.Int, bool) {
// 	var temp byte
// 	xstr := []byte(x.String())
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
// 	flipX.SetString(string(flipStr), 10)
// 	//println(flipX.String())
// 	return flipX, isPalindrome
// }

// /*func IsPalindromic(x *big.Int) (*big.Int, bool) {
// 	divisor.SetInt64(10)
// 	divisorYUp.SetInt64(1)
// 	tempX.Set(x)
// 	flipX.Set(x)
// 	for hold.Div(tempX, divisor).Cmp(bigTen) >= 0 {
// 		divisor.Mul(divisor, bigTen)
// 	}
// 	divisorY.Set(divisor)

// 	isPalindrome := true
// 	for divisor.Cmp(bigTen) >= 0 {
// 		hold.DivMod(tempX, bigTen, checkBegin)
// 		checkEnd.Div(tempX, divisor)
// 		if checkEnd.Cmp(checkBegin) != 0 {
// 			isPalindrome = false
// 			// flipBegin = flipX % (divYUp * 10)
// 			hold.DivMod(flipX, hold.Mul(divisorYUp, bigTen), flipBegin)
// 			// mod = flipBegin % divYUp
// 			hold.DivMod(flipBegin, divisorYUp, mod)
// 			//flipBegin.Sub(flipBegin, mod)
// 			// flipX = flipX - (flipBegin - mod)
// 			flipX.Sub(flipX, flipBegin.Sub(flipBegin, mod))
// 			// flipX = flipX + (checkEnd * divYUp)
// 			flipX.Add(flipX, hold.Mul(divisorYUp, checkEnd))
// 			flipEnd.Div(flipX, divisorY)
// 			hold.DivMod(flipEnd, bigTen, flipEnd)
// 			flipX.Sub(flipX, hold.Mul(flipEnd, divisorY))
// 			flipX.Add(flipX, hold.Mul(checkBegin, divisorY))
// 		}
// 		tempX.Sub(tempX, hold.Mul(divisor, checkEnd))
// 		tempX.Div(tempX, bigTen)
// 		divisor.Div(divisor, bigHundred)
// 		divisorY.Div(divisorY, bigTen)
// 		divisorYUp.Mul(divisorYUp, bigTen)
// 	}
// 	return flipX, isPalindrome
// }*/

// func (l *Lychrel) IsLychrel() bool {
// 	if lychrels[l.x.Int64()] {
// 		return true
// 	}
// 	lyX.Set(l.x)
// 	attemptedBelowMax := []int64{}
// 	if flipX, result = IsPalindromicString(lyX); !result {
// 		attemptedBelowMax = append(attemptedBelowMax, flipX.Int64())
// 	}
// 	lyX.Add(lyX, flipX)
// 	for i := 0; i < MaxTry; i++ {
// 		if lyX.Cmp(bigLyBelowNum) < 0 {
// 			attemptedBelowMax = append(attemptedBelowMax, lyX.Int64())
// 		}
// 		if flipX, result = IsPalindromicString(lyX); result {
// 			return false
// 		}
// 		if lyX.Cmp(bigLyBelowNum) < 0 {
// 			attemptedBelowMax = append(attemptedBelowMax, flipX.Int64())
// 		}
// 		lyX.Add(lyX, flipX)
// 	}
// 	for _, lychrel := range attemptedBelowMax {
// 		lychrels[lychrel] = true
// 	}
// 	if len(lyX.String()) > maxLength {
// 		maxLength = len(lyX.String())
// 	}
// 	return true
// }

// func main() {
// 	l := Lychrel{
// 		x: big.NewInt(32),
// 	}

// 	numLy := 0
// 	var i uint64
// 	for i = 1; i < LyBelowNum; i++ {
// 		l.x.SetUint64(i)
// 		if l.IsLychrel() {
// 			numLy++
// 		}
// 	}
// 	//println(l.IsLychrel())
// 	println(maxLength)
// 	println(numLy)
// }
