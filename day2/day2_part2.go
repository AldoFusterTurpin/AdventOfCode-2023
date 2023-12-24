package main

import (
	"log"
	"strconv"
	"strings"
)

// key: color, value: minimum needed for that color.
type MinimumCubes map[string]int

func (mc MinimumCubes) GetPower() int {
	power := 1
	for _, v := range mc {
		power *= v
	}
	return power
}

func getSumOfPowersOfGames(games []string) int {
	sumOfPowers := 0
	for _, game := range games {
		sumOfPowers += getPowerForGame(game)
	}
	return sumOfPowers
}

func getPowerForGame(game string) int {
	minimumCubesForGame := minimmumSetOfCubesThatMustHaveBeenPresentForGame(game)
	powerForGame := minimumCubesForGame.GetPower()
	return powerForGame
}

func minimmumSetOfCubesThatMustHaveBeenPresentForGame(game string) MinimumCubes {
	minimumCubes := MinimumCubes{"red": 0, "green": 0, "blue": 0}

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

	return minimmumSetOfCubesThatMustHaveBeenPresentForSubsets(subsets, minimumCubes)
}

func minimmumSetOfCubesThatMustHaveBeenPresentForSubsets(subsets []string, minimumCubes MinimumCubes) MinimumCubes {
	// subsets is ["3 blue, 4 red",  "1 red, 2 green, 6 blue", "2 green"]
	for _, s := range subsets {
		updateRequiredMinimumCubesForSubset(s, minimumCubes)
	}
	return minimumCubes
}

func updateRequiredMinimumCubesForSubset(subset string, minimumCubes MinimumCubes) {
	// subset is "3 blue, 4 red"
	subsetList := strings.Split(subset, ", ")
	// ["3 blue", "4 red"]

	for _, s := range subsetList {
		// s -> "3 blue"
		updateMaxForColor(s, minimumCubes)
	}
}

func updateMaxForColor(s string, minimumCubes MinimumCubes) {
	// TrimSpace is important for the first element as it contains a leading whitespace
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
	currentMinimumNeededForColor := minimumCubes[color]
	if n > currentMinimumNeededForColor {
		minimumCubes[color] = n
	}
}
