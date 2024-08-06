package main

func GetTotalWinnings(handsWithBids []HandWithBid) int {
	totalWinnings := 0
	sortHandsWithBidsByStrength(handsWithBids)

	for i, hand := range handsWithBids {
		totalWinnings += hand.bid * (i + 1)
	}
	return totalWinnings
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
