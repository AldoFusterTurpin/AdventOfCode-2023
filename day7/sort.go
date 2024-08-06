package main

import (
	"log"
	"slices"
	"sort"
)

// A hand consists of five cards labeled one of A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, or 2.
// The relative strength of each card follows this order, where A is the highest and 2 is the lowest.
var cardsSortedByStrength = []rune{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}

func sortHandsWithBidsByStrength(hs []HandWithBid) {
	sort.Sort(ByStrength(hs))
}

type ByStrength []HandWithBid

func (b ByStrength) Len() int {
	return len(b)
}

func (b ByStrength) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

// Less implements the function to order the []HandWithBid based on the rank,
// where the weakest hand gets rank 1, the second-weakest hand gets rank 2,
// and so on up to the strongest hand.
// Less should return true when card[i] is weaker than card[j].
func (b ByStrength) Less(i, j int) bool {
	// Refactored (commit: 5f441c351dc8273d725777990eff05402ccdeb31) to use rank integer and comparison
	// for simplicity (compare against previous commit if interested in the diff).
	leftHandRank := b[i].hand.handType.getRank()
	rightHandRank := b[j].hand.handType.getRank()
	if leftHandRank == rightHandRank {
		// When hands have same type, we evaluate the individual cards.
		return b.isLeftHandWeakerThanRightHand(i, j)
	}
	return leftHandRank < rightHandRank
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
		if leftCard != rightCard {
			return b.isLeftCardWeakerThanRightCard(rune(leftCard), rune(rightCard))
		}
	}

	// according to the problem statement, we should never reach this point
	log.Fatalf("error, %v appears twice", b[left])
	return false
}

// isLeftCardWeakerThanRightCard returns true if the individual leftCard is
// weaker than rightCard, otherwise returns false.
func (b ByStrength) isLeftCardWeakerThanRightCard(leftCard, rightCard rune) bool {
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
