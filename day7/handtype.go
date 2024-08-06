package main

import "log"

type HandType string

const (
	FiveOfAKind  HandType = "FiveOfAKind"
	FourOfAkind  HandType = "FourOfAkind"
	FullHouse    HandType = "FullHouse"
	ThreeOfAkind HandType = "ThreeOfAkind"
	TwoPair      HandType = "TwoPair"
	OnePair      HandType = "OnePair"
	HighCard     HandType = "HighCard"
)

// 0 is weakest
// n-1 is strongest
func (h HandType) getRank() int {
	switch h {
	case HighCard:
		return 0
	case OnePair:
		return 1
	case TwoPair:
		return 2
	case ThreeOfAkind:
		return 3
	case FullHouse:
		return 4
	case FourOfAkind:
		return 5
	case FiveOfAKind:
		return 6
	}

	log.Fatal("unexpected error, unknown HandType")
	return -1
}
