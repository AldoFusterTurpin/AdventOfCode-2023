package main

func getFinalNumberOfScratchcards(s string) int {
	cards := convertStringToCards(s)
	nInitialCards := len(cards)

	matchingNumbersMap := getMatchingNumbersMap(cards)
	cardsCounterMap := buildInitialCardsCounter(len(cards))

	// keep total number of scratchcards, will be incremented when passed arround
	totalScratchCards := nInitialCards

	for i := 0; i < nInitialCards; i++ {
		cardId := i + 1
		cardPuntuation := matchingNumbersMap[cardId]
		nCopiesOfThatCard := cardsCounterMap[cardId]

		updateAllCardsWithThatId(cardsCounterMap, cardPuntuation, cardId, &totalScratchCards, nCopiesOfThatCard)
	}
	return totalScratchCards
}

func updateAllCardsWithThatId(cardsCounterMap map[int]int, cardPuntuation int, cardId int, totalScratchCards *int, nCopiesOfThatCard int) {
	for i := 0; i < nCopiesOfThatCard; i++ {
		updateMapOfcardsCounterWithNewCopiesOfSingleCard(cardsCounterMap, cardPuntuation, cardId, totalScratchCards)
	}
}

func updateMapOfcardsCounterWithNewCopiesOfSingleCard(cardsCounterMap map[int]int, cardPuntuation int, cardId int, totalScratchCards *int) {
	nInitialCards := len(cardsCounterMap)
	cardIdToUpdate := cardId + 1

	for i := 0; i < cardPuntuation; i++ {
		// "Cards will never make you copy a card past the end of the table"
		if cardIdToUpdate > nInitialCards {
			return
		}

		*totalScratchCards++
		cardsCounterMap[cardIdToUpdate] += 1

		// Uncomment line below and run the tests (in verbose mode, go test -v ./...) or execute the main with the full input.
		// You will see execution time go from just < 5 seconds to > 30 seconds.
		// I/O is expensive 👀
		// log.Printf("totalScratchCards so far: %v\n", *totalScratchCards)

		cardIdToUpdate++
	}
}

// buildInitialCardsCounter() return a map where key is the Card id (starting at i, not 0)
// and value is how many cards we have in total, including the copies.
func buildInitialCardsCounter(nCards int) map[int]int {
	m := make(map[int]int)
	for i := 0; i < nCards; i++ {
		m[i+1] = 1
	}

	return m
}

// getMatchingNumbersMap returns a map where key is cardId
// and value how many matching numbers has the card with that cardId
func getMatchingNumbersMap(cards []Card) map[int]int {
	m := make(map[int]int)
	for _, card := range cards {
		m[card.id] = card.getMatchingNumbers()
	}

	return m
}
