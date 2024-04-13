package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

// A hand consists of five cards labeled one of A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, or 2.
// The relative strength of each card follows this order, where A is the highest and 2 is the lowest.
var cardsSortedByStrength = []rune{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}

type HandType string

const (
	FiveOfAKind  = "FiveOfAKind"
	FourOfAkind  = "FourOfAkind"
	FullHouse    = "FullHouse"
	ThreeOfAkind = "ThreeOfAkind"
	TwoPair      = "TwoPair"
	OnePair      = "OnePair"
	HighCard     = "HighCard"
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

func sortByStrength(hs []HandWithBid) {
	sort.Sort(ByStrength(hs))
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

type ByStrength []HandWithBid

func (o ByStrength) Len() int {
	return len(o)
}

func (o ByStrength) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}

// Less implements the function to order the []HandWithBid based on the rank,
// where the weakest hand gets rank 1, the second-weakest hand gets rank 2,
// and so on up to the strongest hand. So Less should return true when card[i] is weaker than card[j]
func (b ByStrength) Less(left, right int) bool {
	// Return true when left is weaker than right.
	// Refactored to use rank integer and comparison for simplicity (compare against previous commit if interested in the diff)
	leftHandRank := b[left].hand.handType.getRank()
	rightHandRank := b[right].hand.handType.getRank()
	if leftHandRank == rightHandRank {
		return b.isLeftHandWeakerThanRightHand(left, right)
	}
	return leftHandRank > rightHandRank

}

func (b ByStrength) isLeftHandWeakerThanRightHand(left, right int) bool {
	// If two hands have the same type, a second ordering rule takes effect.
	// Start by comparing the first card in each hand.
	// If these cards are different, the hand with the stronger first card is considered stronger.
	// If the first card in each hand have the same label, however, then move on to considering the second card in each hand.

	nCardsLeft := len(b[left].hand.cards)
	nCardsRight := len(b[right].hand.cards)

	// this should never happen
	if nCardsLeft != nCardsRight {
		log.Fatal("error, not the same number of cards in left hand and right hand")
	}

	for index := 0; index < nCardsLeft; index++ {
		leftCard := b[left].hand.cards[index]
		rightCard := b[right].hand.cards[index]
		if leftCard == rightCard {
			continue
		}
		return isLeftCardWeakerThanRightCard(rune(leftCard), rune(rightCard))
	}

	// according to the problem statement, we should never reach this point
	log.Fatalf("error, %v appears twice", b[left])
	return false
}

// isLeftCardWeakerThanRightCard returns true if leftCard is stronger than rightCard, otherwise returns false.
func isLeftCardWeakerThanRightCard(leftCard, rightCard rune) bool {
	leftCardIndex := slices.Index(cardsSortedByStrength, leftCard)
	if leftCardIndex == -1 {
		log.Fatalf("leftCard %v is not present in cardsSortedByStrength", string(leftCard))
	}

	rightCardIndex := slices.Index(cardsSortedByStrength, rightCard)
	if rightCardIndex == -1 {
		log.Fatalf("rightCard %v is not present in cardsSortedByStrength", string(rightCard))
	}

	return leftCardIndex > rightCardIndex
}

func main() {
	folder := "input_files"
	fileName := "sample"
	fileExtension := ".txt"

	fileContentBytes, err := os.ReadFile(folder + "/" + fileName + fileExtension)
	if err != nil {
		log.Fatal(err)
	}

	fileContent := string(fileContentBytes)

	handsWithBids, err := processFileContent(fileContent)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("handsWithBids: %+v", handsWithBids)
	// fmt.Println("result is", result)
}

func processFileContent(fileContent string) ([]HandWithBid, error) {
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
