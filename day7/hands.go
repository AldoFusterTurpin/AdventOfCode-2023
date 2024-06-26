package main

import (
	"strconv"
	"strings"
)

// A hand consists of five cards labeled one of A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, or 2.
// The relative strength of each card follows this order, where A is the highest and 2 is the lowest.
var cardsSortedByStrength = []rune{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}

func GetTotalWinnings(hs []HandWithBid) int {
	totalWinnings := 0
	sortHandsWithBidsByStrength(hs)
	for i, h := range hs {
		totalWinnings += h.bid * (i + 1)
	}
	return totalWinnings
}

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

// 0 is strongest
// n -1 is weakest
func (h HandType) getRank() int {
	switch h {
	case FiveOfAKind:
		return 0
	case FourOfAkind:
		return 1
	case FullHouse:
		return 2
	case ThreeOfAkind:
		return 3
	case TwoPair:
		return 4
	case OnePair:
		return 5
	case HighCard:
		return 6
	}
	return -1
}

type Hand struct {
	cards    string
	handType HandType
}

func (h *Hand) setHandType() {
	// key: letter
	// value: how many times it appears in cards
	m := make(map[rune]int)

	for _, r := range h.cards {
		m[r]++
	}
	if len(m) == 1 {
		h.handType = FiveOfAKind
		return
	}
	if len(m) == 4 {
		h.handType = OnePair
		return
	}
	if len(m) == 5 {
		h.handType = HighCard
		return
	}
	for _, v := range m {
		if v == 4 {
			h.handType = FourOfAkind
			return
		}
		// where three cards have the same label
		if v == 3 {
			// and the remaining two cards share a different label
			if len(m) == 2 {
				h.handType = FullHouse
				return
			}
			if len(m) == 3 {
				h.handType = ThreeOfAkind
				return
			}
		}
		if v == 2 {
			if len(m) == 3 {
				h.handType = TwoPair
				return
			}
		}
	}
}

func NewHand(cards string) *Hand {
	h := &Hand{cards: cards}
	h.setHandType()
	return h
}

type HandWithBid struct {
	hand Hand
	bid  int
}

func NewHandWithBid(cards, bid string) (*HandWithBid, error) {
	bidInt, err := strconv.Atoi(bid)
	if err != nil {
		return nil, err
	}

	handWithBid := &HandWithBid{
		hand: *NewHand(cards),
		bid:  bidInt,
	}

	return handWithBid, nil
}

func getHandsWithBidsFromFileContent(fileContent string) ([]HandWithBid, error) {
	var handsWithBids []HandWithBid
	lines := strings.Split(fileContent, "\n")
	for _, line := range lines {
		line = strings.Trim(line, " ")
		lineSplitted := strings.Split(line, " ")

		cards := lineSplitted[0]
		bid := lineSplitted[1]
		handWithBid, err := NewHandWithBid(cards, bid)
		if err != nil {
			return nil, err
		}
		handsWithBids = append(handsWithBids, *handWithBid)
	}

	return handsWithBids, nil
}
