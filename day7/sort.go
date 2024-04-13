package main

import (
	"log"
	"slices"
	"sort"
)

func sortByStrength(hs []HandWithBid) {
	sort.Sort(ByStrength(hs))
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
		return b.isLeftCardWeakerThanRightCard(rune(leftCard), rune(rightCard))
	}

	// according to the problem statement, we should never reach this point
	log.Fatalf("error, %v appears twice", b[left])
	return false
}

// isLeftCardWeakerThanRightCard returns true if leftCard is stronger than rightCard, otherwise returns false.
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
