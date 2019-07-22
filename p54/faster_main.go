package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"strings"
// )

// const (
// 	handSize = 5
// )

// type Suit int

// const (
// 	Heart Suit = iota
// 	Club
// 	Spade
// 	Diamond
// 	NoSuit
// )

// type Card struct {
// 	Suit  Suit
// 	Value int
// }

// var suitValues = map[string]Suit{
// 	"H": Heart,
// 	"S": Spade,
// 	"D": Diamond,
// 	"C": Club,
// }

// var cardValues = map[string]int{
// 	"2": 1,
// 	"3": 2,
// 	"4": 3,
// 	"5": 4,
// 	"6": 5,
// 	"7": 6,
// 	"8": 7,
// 	"9": 8,
// 	"T": 9,
// 	"J": 10,
// 	"Q": 11,
// 	"K": 12,
// 	"A": 13,
// }

// // First return val is hand rank, second is highest card in hand, third is ordered list of high to low
// func analyzeHand(pairs map[int]int, order []int, flush bool) (int, int, []int) {
// 	straight := false
// 	if len(pairs) == handSize && order[0]-order[handSize-1] == handSize-1 {
// 		straight = true
// 	}

// 	if flush && straight {
// 		return 10, order[0], order
// 	} else if len(pairs) == 2 {
// 		for k, v := range pairs {
// 			if v == 4 {
// 				return 9, k, order
// 			} else if v == 3 {
// 				return 8, k, order
// 			}
// 		}
// 	} else if flush {
// 		return 7, order[0], order
// 	} else if straight {
// 		return 6, order[0], order
// 	} else if len(pairs) == 3 {
// 		highTwoPair := 0
// 		for k, v := range pairs {
// 			if v == 3 {
// 				return 5, order[0], order
// 			} else if v == 2 {
// 				if k > highTwoPair {
// 					highTwoPair = k
// 				}
// 			}
// 		}
// 		return 4, highTwoPair, order
// 	} else if len(pairs) == 4 {
// 		for k, v := range pairs {
// 			if v == 2 {
// 				return 3, k, order
// 			}
// 		}
// 	}

// 	return 2, order[0], order
// }

// func main() {
// 	dat, err := ioutil.ReadFile("./poker.txt")
// 	if err != nil {
// 		fmt.Printf("Couldn't read file yo: %s\n", err.Error())
// 		return
// 	}

// 	playerOneWins := 0
// 	lines := strings.Split(string(dat), "\n")

// 	for _, line := range lines {
// 		pairs1 := map[int]int{}
// 		order1 := make([]int, handSize)
// 		flush1 := true
// 		suit1 := NoSuit
// 		pairs2 := map[int]int{}
// 		order2 := make([]int, handSize)
// 		flush2 := true
// 		suit2 := NoSuit
// 		for i, card := range strings.Split(line, " ") {
// 			cardSuit := suitValues[string(card[1])]
// 			cardValue := cardValues[string(card[0])]
// 			if i < handSize {
// 				pairs1[cardValue]++
// 				if suit1 == NoSuit {
// 					suit1 = cardSuit
// 				} else if suit1 != cardSuit {
// 					flush1 = false
// 				}
// 				order1[i] = cardValue
// 				for j := i; j > 0; j-- {
// 					if order1[j] > order1[j-1] {
// 						tmp := order1[j]
// 						order1[j] = order1[j-1]
// 						order1[j-1] = tmp
// 					}
// 				}
// 			} else {
// 				pairs2[cardValue]++
// 				if suit2 == NoSuit {
// 					suit2 = cardSuit
// 				} else if suit2 != cardSuit {
// 					flush2 = false
// 				}
// 				order2[i-handSize] = cardValue
// 				for j := i - handSize; j > 0; j-- {
// 					if order2[j] > order2[j-1] {
// 						tmp := order2[j]
// 						order2[j] = order2[j-1]
// 						order2[j-1] = tmp
// 					}
// 				}
// 			}
// 		}

// 		h1Score, h1High, h1Order := analyzeHand(pairs1, order1, flush1)
// 		h2Score, h2High, h2Order := analyzeHand(pairs2, order2, flush2)

// 		if h1Score > h2Score {
// 			playerOneWins++
// 		} else if h2Score == h1Score {
// 			if h1High > h2High {
// 				playerOneWins++
// 			} else if h1High == h2High {
// 				for i, val := range h1Order {
// 					if val > h2Order[i] {
// 						playerOneWins++
// 						break
// 					} else if h2Order[i] > val {
// 						break
// 					}
// 				}
// 			}
// 		}
// 	}

// 	fmt.Printf("And the result is..... %d\n", playerOneWins)
// }
