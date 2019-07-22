package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Suit int

const (
	Heart Suit = iota
	Club
	Spade
	Diamond
	NoSuit

	handSize = 5
)

type Card struct {
	Suit  Suit
	Value int
}

var suitValues = map[string]Suit{
	"H": Heart,
	"S": Spade,
	"D": Diamond,
	"C": Club,
}

var cardValues = map[string]int{
	"2": 1,
	"3": 2,
	"4": 3,
	"5": 4,
	"6": 5,
	"7": 6,
	"8": 7,
	"9": 8,
	"T": 9,
	"J": 10,
	"Q": 11,
	"K": 12,
	"A": 13,
}

// First return val is hand rank, second is highest card in hand, third is ordered list of high to low
func analyzeHand(hand []Card) (int, int) {
	pairs := map[int]int{}
	flush := true
	suit := NoSuit
	highCard := 0
	lowCard := 1000

	for _, card := range hand {
		pairs[card.Value]++
		if suit == NoSuit {
			suit = card.Suit
		} else if suit != card.Suit {
			flush = false
		}
		if highCard < card.Value {
			highCard = card.Value
		}
		if lowCard > card.Value {
			lowCard = card.Value
		}
	}
	straight := false
	if len(pairs) == handSize && highCard-lowCard == handSize-1 {
		straight = true
	}

	if flush && straight {
		return 10, highCard
	} else if len(pairs) == 2 {
		for k, v := range pairs {
			if v == 4 {
				return 9, k
			} else if v == 3 {
				return 8, k
			}
		}
		fmt.Println("something has gone horribly wrong in 4ofkind or full house")
		fmt.Printf("Hand: %#v\n", hand)
	} else if flush {
		return 7, highCard
	} else if straight {
		return 6, highCard
	} else if len(pairs) == 3 {
		highTwoPair := 0
		for k, v := range pairs {
			if v == 3 {
				return 5, highCard
			} else if v == 2 {
				if k > highTwoPair {
					highTwoPair = k
				}
			}
		}
		return 4, highTwoPair
	} else if len(pairs) == 4 {
		for k, v := range pairs {
			if v == 2 {
				return 3, k
			}
		}
		fmt.Println("something has gone horribly wrong in pair")
		fmt.Printf("Hand: %#v\n", hand)
	}

	return 2, highCard
}

func main() {
	dat, err := ioutil.ReadFile("./poker.txt")
	if err != nil {
		fmt.Printf("Couldn't read file yo: %s\n", err.Error())
		return
	}

	playerOneWins := 0
	lines := strings.Split(string(dat), "\n")

	playerOneHand := make([]Card, 5)
	playerTwoHand := make([]Card, 5)
	for _, line := range lines {
		for i, card := range strings.Split(line, " ") {
			if i < handSize {
				playerOneHand[i] = Card{Suit: suitValues[string(card[1])], Value: cardValues[string(card[0])]}
			} else {
				playerTwoHand[i-handSize] = Card{Suit: suitValues[string(card[1])], Value: cardValues[string(card[0])]}
			}
		}

		h1Score, h1High := analyzeHand(playerOneHand)
		h2Score, h2High := analyzeHand(playerTwoHand)

		if h1Score > h2Score {
			playerOneWins++
		} else if h2Score == h1Score {
			if h1High > h2High {
				playerOneWins++
			} else if h1High == h2High {
				h1Order := make([]int, 5)
				h2Order := make([]int, 5)
				for i, card := range playerOneHand {
					h1Order[i] = card.Value
					for j := i; j > 0; j-- {
						if h1Order[j] > h1Order[j-1] {
							tmp := h1Order[j]
							h1Order[j] = h1Order[j-1]
							h1Order[j-1] = tmp
						}
					}
				}
				for i, card := range playerTwoHand {
					h2Order[i] = card.Value
					for j := i; j > 0; j-- {
						if h2Order[j] > h2Order[j-1] {
							tmp := h2Order[j]
							h2Order[j] = h2Order[j-1]
							h2Order[j-1] = tmp
						}
					}
				}
				for i, val := range h1Order {
					if val > h2Order[i] {
						playerOneWins++
						break
					} else if h2Order[i] > val {
						break
					}
				}
			}
		}
	}

	fmt.Printf("And the result is..... %d\n", playerOneWins)
}
