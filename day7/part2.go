package main

import (
	"log"
	"slices"
	"sort"
	"strconv"
	"strings"
)

// A is the strongest. J is the weakest.
var cardsSortedByStrengthWithJoker = []rune{'A', 'K', 'Q', 'T', '9', '8', '7', '6', '5', '4', '3', '2', 'J'}

const jokerString = "J"
const jokerRune = 'J'

func GetTotalWinningsPart2(handsWithBids []HandWithBid) int {
	totalWinnings := 0
	sortPart2(handsWithBids)

	for i, handWithBid := range handsWithBids {
		totalWinnings += handWithBid.bid * (i + 1)
	}
	return totalWinnings
}

func getHandsWithBidsFromFileContentPart2(fileContent string) ([]HandWithBid, error) {
	var handsWithBids []HandWithBid
	lines := strings.Split(fileContent, "\n")
	for _, line := range lines {
		line = strings.Trim(line, " ")
		lineSplitted := strings.Split(line, " ")

		cards := lineSplitted[0]
		bid := lineSplitted[1]
		handWithBid, err := NewHandWithBidPart2(cards, bid)
		if err != nil {
			return nil, err
		}
		handsWithBids = append(handsWithBids, *handWithBid)
	}

	return handsWithBids, nil
}

func NewHandWithBidPart2(cards, bid string) (*HandWithBid, error) {
	bidInt, err := strconv.Atoi(bid)
	if err != nil {
		return nil, err
	}

	handWithBid := &HandWithBid{
		hand: *NewHandPart2(cards),
		bid:  bidInt,
	}

	return handWithBid, nil
}

func NewHandPart2(cards string) *Hand {
	h := &Hand{cards: cards}

	jokerIndex := strings.Index(cards, jokerString)
	noJokerPresent := jokerIndex == -1
	if noJokerPresent {
		h.setHandType()
		return h
	}
	h.handType = getHandTypeWhenJokerIsPresent(cards)
	return h
}

func getHandTypeWhenJokerIsPresent(cards string) HandType {
	indexesOfJokers, _ := getIndexes(cards, jokerRune)
	nJokers := len(indexesOfJokers)
	uniquesWithoutJoker := getCharCounterWithoutJoker(cards)
	nUniqueNumbersWithoutJoker := len(uniquesWithoutJoker)

	if nJokers == 5 || nJokers == 4 {
		// JJJJJ  -> JJJJJ
		// JJJJ 2 -> JJJJJ
		return FiveOfAKind
	}

	if nJokers == 3 {
		if nUniqueNumbersWithoutJoker == 2 {
			// JJJ 91 to -> 9999 1
			return FourOfAkind
		}
		// else if nUniqueNumbersWithoutJoker == 1 ...
		// From JJJ 99 to -> 999 99
		return FiveOfAKind
	}

	if nJokers == 2 {
		if nUniqueNumbersWithoutJoker == 1 {
			// From JJ 999 to -> 999 99
			return FiveOfAKind
		}

		// if two remaining cards are equal
		if nUniqueNumbersWithoutJoker == 2 {
			// JJ 99 7 to -> 9999 7
			return FourOfAkind
		}

		// if three remaining cards are different
		if nUniqueNumbersWithoutJoker == 3 {
			// From JJ 987 to -> 999 87
			return ThreeOfAkind
		}
	}

	if nJokers == 1 {
		if nUniqueNumbersWithoutJoker == 4 {
			// From J 9876 to -> 99 876
			return OnePair
		}
		//TODO: Aldo
		if nUniqueNumbersWithoutJoker == 3 {
			// From J 99 76 to -> 999 76
			return ThreeOfAkind
		}

		if nUniqueNumbersWithoutJoker == 2 {
			// J 22 33 -> 22 333 full house.
			// uniquesWithoutJoler would be -> '2':2, '3':2
			// We have FullHouse when all the values of all the keys in uniquesWithoutJoler
			// have the same value. Otherwise we have a FourOfAKind
			first := true
			firstValue := 0
			for _, v := range uniquesWithoutJoker {
				if first {
					firstValue = v
					first = false
				} else {
					// From J 222 3 -> 222 33 four of a kind
					// uniquesWithoutJoler would be -> '2':3, '3':1 -> not all the keys have the same value.
					if v != firstValue {
						return FourOfAkind
					}
				}
			}
			return FullHouse
		}

		if nUniqueNumbersWithoutJoker == 1 {
			// From J 9999 to ->  99999
			return FiveOfAKind
		}
	}

	log.Fatal("unexpected error, wrong number of joker and no joker cards")
	return FullHouse //no matter what we return, impossible to reach
}

// getCharCounterWithoutJoker returns the unique runes contained in s without taking into account the jokerRune
// and how many times it appears. key: rune (letter), value: number of times it appeats in s
func getCharCounterWithoutJoker(s string) map[rune]int {
	m := make(map[rune]int)
	for _, v := range s {
		if v != jokerRune {
			m[v]++
		}
	}
	return m
}

// getIndexes iterates over s and put all the indexes of runes that are not the targetChar in
// indexesOfNoJokers and it puts all the indexes of runes == targetChar in indexesOfJokers
func getIndexes(cards string, targetChar rune) (indexesOfTarget []int, indexesOfNoTarget []int) {
	for index, card := range cards {
		if card == targetChar {
			indexesOfTarget = append(indexesOfTarget, index)
		} else {
			indexesOfNoTarget = append(indexesOfNoTarget, index)
		}
	}
	return indexesOfTarget, indexesOfNoTarget
}

func sortPart2(handsWithBids []HandWithBid) {
	sort.Sort(ByStrengthPart2(handsWithBids))
}

type ByStrengthPart2 []HandWithBid

func (b ByStrengthPart2) Len() int {
	return len(b)
}

func (b ByStrengthPart2) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

// Less implements the function to order the []HandWithBid based on the rank,
// where the weakest hand gets rank 1, the second-weakest hand gets rank 2,
// and so on up to the strongest hand.
// Less should return true when card[i] is weaker than card[j].
func (b ByStrengthPart2) Less(i, j int) bool {
	leftHand := b[i].hand
	leftHandRank := leftHand.handType.getRank()
	rightHand := b[j].hand
	rightHandRank := rightHand.handType.getRank()

	if leftHandRank != rightHandRank {
		return leftHandRank < rightHandRank
	}

	jokerPresentInLeftHand := strings.Contains(leftHand.cards, jokerString)
	jokerPresentInRighttHand := strings.Contains(rightHand.cards, jokerString)
	if jokerPresentInLeftHand || jokerPresentInRighttHand {
		return b.isLeftHandWeakerThanRightHandWhenJoker(i, j)
	}
	return b.isLeftHandWeakerThanRightHandWhenNoJoker(i, j)
}

func (b ByStrengthPart2) isLeftHandWeakerThanRightHandWhenNoJoker(left, right int) bool {
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

func (b ByStrengthPart2) isLeftHandWeakerThanRightHandWhenJoker(leftIndex, rightIndex int) bool {
	nCardsLeft := len(b[leftIndex].hand.cards)
	nCardsRight := len(b[rightIndex].hand.cards)
	// this should never happen
	if nCardsLeft != nCardsRight {
		log.Fatal("unexpected error, not the same number of cards in left and right hand")
	}

	for index := 0; index < nCardsLeft; index++ {
		leftCard := b[leftIndex].hand.cards[index]
		rightCard := b[rightIndex].hand.cards[index]
		if leftCard != rightCard {
			return b.isLeftCardWeakerThanRightCardWhenJoker(rune(leftCard), rune(rightCard))
		}
	}

	// according to the problem statement, we should never reach this point
	log.Fatalf("error, %v appears twice", b[leftIndex])
	return false
}

func (b ByStrengthPart2) isLeftCardWeakerThanRightCard(leftCard, rightCard rune) bool {
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

// isLeftCardWeakerThanRightCardWhenJoker should return true when the card with leftCard index
// is weaker than the cards with rightCard index.
func (b ByStrengthPart2) isLeftCardWeakerThanRightCardWhenJoker(leftCard, rightCard rune) bool {
	leftCardIndex := slices.Index(cardsSortedByStrengthWithJoker, leftCard)
	if leftCardIndex == -1 {
		log.Fatalf("leftCard %v is not present in cardsSortedByStrengthWithJoker", string(leftCard))
	}

	rightCardIndex := slices.Index(cardsSortedByStrengthWithJoker, rightCard)
	if rightCardIndex == -1 {
		log.Fatalf("rightCard %v is not present in cardsSortedByStrengthWithJoker", string(rightCard))
	}

	return leftCardIndex > rightCardIndex
}
