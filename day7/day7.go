package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// Part2: 252898370 correct answer

const part2 = true

func main() {
	folder := "input_files"
	fileName := "input"
	fileExtension := ".txt"

	fileContentBytes, err := os.ReadFile(folder + "/" + fileName + fileExtension)
	if err != nil {
		log.Fatal(err)
	}

	inputString := string(fileContentBytes)
	inputString = strings.TrimSpace(inputString)

	result := getResult(inputString)
	fmt.Println("result is", result)
}

func getResult(s string) int {
	if part2 {
		hs, err := getHandsWithBidsFromFileContentPart2(s)
		if err != nil {
			log.Fatal(err)
		}
		return GetTotalWinningsPart2(hs)
	}

	hs, err := getHandsWithBidsFromFileContent(s)
	if err != nil {
		log.Fatal(err)
	}
	return GetTotalWinnings(hs)
}
