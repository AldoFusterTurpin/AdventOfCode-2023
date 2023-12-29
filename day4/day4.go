package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	multipleSpacesRegex = regexp.MustCompile(`\s+`)
)

func main() {
	folder := "input_files"
	fileName := "input"
	fileExtension := ".txt"
	fileContent, err := os.ReadFile(folder + "/" + fileName + fileExtension)
	if err != nil {
		log.Fatal(err)
	}

	s := string(fileContent)
	s = strings.TrimSpace(s)

	// sum := getTotalScratchcardsPuntuation(s) // part 1

	sum := getFinalNumberOfScratchcards(s) // part 2

	fmt.Printf("result %v ", sum)
}

type Card struct {
	id             int
	winningNumbers []int
	numbersWeHave  []int
}

func convertStringToCards(s string) []Card {
	var cards []Card
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		line = multipleSpacesRegex.ReplaceAllString(line, " ")
		line = strings.TrimSpace(line)

		// "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
		slice := strings.Split(line, ":")
		// ["Card 1",  "41 48 83 86 17 | 83 86  6 31 17  9 48 53"]

		cardString := slice[0]
		// "Card 1"

		cardStringSplitted := strings.Split(cardString, " ")
		// ["Car", "1"]

		cardIdStr := cardStringSplitted[1]
		// "1"

		cardId, err := strconv.Atoi(cardIdStr)
		if err != nil {
			log.Fatal(err)
		}

		rightPart := slice[1]
		// rightPart is "41 48 83 86 17 | 83 86  6 31 17  9 48 53"

		rightPartSplitted := strings.Split(rightPart, "|")
		// rightPartSplitted ["41 48 83 86 17", "83 86  6 31 17  9 48 53"]

		winningNumbersStr := rightPartSplitted[0]
		// "41 48 83 86 17"
		winningNumbers := convertStringToSliceOfInts(winningNumbersStr)

		numbersWeHaveStr := rightPartSplitted[1]
		// "83 86  6 31 17  9 48 53"
		numbersWeHave := convertStringToSliceOfInts(numbersWeHaveStr)

		c := Card{
			id:             cardId,
			numbersWeHave:  numbersWeHave,
			winningNumbers: winningNumbers,
		}
		cards = append(cards, c)
	}

	return cards
}

func getTotalScratchcardsPuntuation(s string) int {
	cards := convertStringToCards(s)
	sum := 0

	for _, card := range cards {
		sum += card.getPuntuation()
	}

	return sum
}

func (card Card) getPuntuation() int {
	firstMatch := true
	puntuation := 0

	winningNumbersMap := getUniqueWinningNumbers(card.winningNumbers)

	for _, v := range card.numbersWeHave {
		if winningNumbersMap[v] {
			if firstMatch {
				puntuation = 1
			} else {
				puntuation *= 2
			}
			firstMatch = false
		}
	}

	return puntuation
}

func (card Card) getMatchingNumbers() int {
	n := 0
	winningNumbersMap := getUniqueWinningNumbers(card.winningNumbers)

	for _, v := range card.numbersWeHave {
		if winningNumbersMap[v] {
			n++
		}
	}

	return n
}

func getUniqueWinningNumbers(allWinningNumbers []int) map[int]bool {
	uniqueWinningNumbers := make(map[int]bool)
	for _, v := range allWinningNumbers {
		uniqueWinningNumbers[v] = true
	}
	return uniqueWinningNumbers
}

func convertStringToSliceOfInts(numbers string) []int {
	numbers = strings.TrimSpace(numbers)
	var result []int

	numbersSLiceStr := strings.Split(numbers, " ")
	for _, s := range numbersSLiceStr {
		v, err := strconv.Atoi(s)
		if err != nil {
			// for simplification...
			log.Fatal(err)
		}

		result = append(result, v)
	}
	return result
}
