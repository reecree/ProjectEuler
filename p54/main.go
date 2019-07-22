package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"strings"
// )

// type Suit int

// const (
// 	Heart Suit = iota
// 	Club
// 	Spade
// 	Diamond
// 	NoSuit

// 	handSize = 5
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
// func analyzeHand(hand []Card) (int, int, []int) {
// 	pairs := map[int]int{}
// 	order := make([]int, handSize)
// 	flush := true
// 	suit := NoSuit

// 	for i, card := range hand {
// 		pairs[card.Value]++
// 		if suit == NoSuit {
// 			suit = card.Suit
// 		} else if suit != card.Suit {
// 			flush = false
// 		}
// 		order[i] = card.Value
// 		for j := i; j > 0; j-- {
// 			if order[j] > order[j-1] {
// 				tmp := order[j]
// 				order[j] = order[j-1]
// 				order[j-1] = tmp
// 			}
// 		}
// 	}
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
// 		fmt.Println("something has gone horribly wrong in 4ofkind or full house")
// 		fmt.Printf("Hand: %#v\n", hand)
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
// 		fmt.Println("something has gone horribly wrong in pair")
// 		fmt.Printf("Hand: %#v\n", hand)
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

// 	playerOneHand := make([]Card, 5)
// 	playerTwoHand := make([]Card, 5)
// 	for _, line := range lines {
// 		for i, card := range strings.Split(line, " ") {
// 			if i < handSize {
// 				//checkList1[cardValues[string(card[0])]] = true
// 				playerOneHand[i] = Card{Suit: suitValues[string(card[1])], Value: cardValues[string(card[0])]}
// 			} else {
// 				//checkList2[cardValues[string(card[0])]] = true
// 				playerTwoHand[i-handSize] = Card{Suit: suitValues[string(card[1])], Value: cardValues[string(card[0])]}
// 			}
// 		}

// 		h1Score, h1High, h1Order := analyzeHand(playerOneHand)
// 		h2Score, h2High, h2Order := analyzeHand(playerTwoHand)

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
