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
	fileName := "input"
	fileExtension := ".txt"
	fileContent, err := os.ReadFile(folder + "/" + fileName + fileExtension)
	if err != nil {
		log.Fatal(err)
	}

	s := string(fileContent)
	s = strings.TrimSpace(s)
	lines := getLinesFromFileContent(s)
	// fmt.Println(lines)

	// sum := sumOfIDsOfThePossibleGames(lines) // Part 1
	sum := getSumOfPowersOfGames(lines) // Part 2
	fmt.Printf("sum is %v: ", sum)
}

func getLinesFromFileContent(s string) []string {
	return strings.Split(s, "\n")
}

func sumOfIDsOfThePossibleGames(games []string) int {
	sum := 0
	for i, game := range games {
		gameID := i + 1
		if isGamePossible(game) {
			sum += gameID
		}
	}
	return sum
}

func isGamePossible(game string) bool {
	// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	game = strings.TrimSpace(game)

	gameSlice := strings.Split(game, ":")
	// ["Game 1", "3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"]

	// gameId := gameSlice[0]
	//"Game 1"

	subsetsStr := gameSlice[1]
	// "3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"

	subsets := strings.Split(subsetsStr, ";")
	// ["3 blue, 4 red",  "1 red, 2 green, 6 blue", "2 green"]

	return areSubsetsPossible(subsets)
}

// areSubsetsPossible receives subsets like
// ["3 blue, 4 red",  "1 red, 2 green, 6 blue", "2 green"]
// and returns if it is possible according to the game rules.
func areSubsetsPossible(subsets []string) bool {
	for _, s := range subsets {
		if !isSubsetPossible(s) {
			return false
		}
	}
	return true
}

// isSubsetPossible receives a subset like
// "3 blue, 4 red" and returns if it is possible
// according to the game rules.
func isSubsetPossible(subset string) bool {
	subsetList := strings.Split(subset, ", ")
	// ["3 blue", "4 red"]

	for _, s := range subsetList {
		// s -> "3 blue"
		if !isSingleGuessPossible(s) {
			return false
		}
	}
	return true
}

// isSingleGuessPossible receives  a string like
// "3 blue" or " 3 blue" and returns if it is possible according to the game rules.
func isSingleGuessPossible(s string) bool {
	// TrimSpace is important for the first guess as it contains a leading whitespace
	// " 3 blue"
	s = strings.TrimSpace(s)
	// "3 blue"

	l := strings.Split(s, " ")
	// ["3", "blue"]

	nStr := l[0]
	// "3"

	n, err := strconv.Atoi(nStr)
	// 3
	if err != nil {
		log.Fatal(err)
	}

	color := l[1]
	// "blue"
	return n <= maxNumberOfCubesPossibleForThatColor(color)
}

func maxNumberOfCubesPossibleForThatColor(s string) int {
	m := map[string]int{"red": 12, "green": 13, "blue": 14}
	return m[s]
}
