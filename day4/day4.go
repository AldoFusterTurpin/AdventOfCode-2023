package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	folder := "input_files"
	fileName := "sample"
	fileExtension := ".txt"
	fileContent, err := os.ReadFile(folder + "/" + fileName + fileExtension)
	if err != nil {
		log.Fatal(err)
	}

	s := string(fileContent)
	s = strings.TrimSpace(s)

	// sum := getSumOfPartNumbers(s) // part 1
	sum := getTotalScratchcardsPuntuation(s)
	fmt.Printf("sum of gear ratios is: %v ", sum)
}

func getTotalScratchcardsPuntuation(s string) int {
	sum := 0
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		sum += getCardPuntuation(line)
	}
	return sum
}

func getCardPuntuation(line string) int {
	// "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
	slice := strings.Split(line, ":")
	// ["Card 1",  "41 48 83 86 17 | 83 86  6 31 17  9 48 53"]

	rightPart := slice[1]
	// rightPart is "41 48 83 86 17 | 83 86  6 31 17  9 48 53"

	rightPartSplitted := strings.Split(rightPart, "|")
	// rightPartSplitted ["41 48 83 86 17", "83 86  6 31 17  9 48 53"]

	winningNumbers := rightPartSplitted[0]
	// "41 48 83 86 17"

	numbersWeHave := rightPartSplitted[1]
	// "83 86  6 31 17  9 48 53"

	winningNumbersMap := mapFromWinningNumbers(winningNumbers)
	return getPuntuationForCard(numbersWeHave, winningNumbersMap)
}

// mapFromWinningNumbers receives an string like "41 48 83 86 17" and return a map (used as a set)
// of the unique numbers.
func mapFromWinningNumbers(winningNumbers string) map[int]bool {
	m := make(map[int]bool)

	winningNumbers = strings.TrimSpace(winningNumbers)
	winningNumbersSlice := strings.Split(winningNumbers, " ")
	for _, s := range winningNumbersSlice {
		// some numbers include more than one space,
		// we need to check for them and skip
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}

		v, err := strconv.Atoi(s)
		if err != nil {
			// In general, I would return the error and handle it outside but I assume clean input here for simplification.
			log.Fatal(err)
		}

		m[v] = true
	}
	return m
}

func getPuntuationForCard(numbersWeHaveStr string, winningNumbersMap map[int]bool) int {
	numbersWeHaveStr = strings.TrimSpace(numbersWeHaveStr)
	numbersWeHave := convertStringToSliceOfInts(numbersWeHaveStr)
	firstMatch := true
	puntuation := 0

	for _, v := range numbersWeHave {
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

func convertStringToSliceOfInts(numbers string) []int {
	numbers = strings.TrimSpace(numbers)
	var result []int

	numbersSLiceStr := strings.Split(numbers, " ")
	for _, s := range numbersSLiceStr {
		// some numbers include more than one space,
		// we need to check for them and skip
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}

		v, err := strconv.Atoi(s)
		if err != nil {
			// In general, I would return the error and handle it outside but I assume clean input here for simplification.
			log.Fatal(err)
		}

		result = append(result, v)
	}
	return result
}
