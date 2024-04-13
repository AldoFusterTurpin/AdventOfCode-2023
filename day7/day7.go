package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	folder := "input_files"
	fileName := "sample"
	fileExtension := ".txt"

	fileContentBytes, err := os.ReadFile(folder + "/" + fileName + fileExtension)
	if err != nil {
		log.Fatal(err)
	}

	fileContent := string(fileContentBytes)

	handsWithBids, err := processFileContent(fileContent)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("handsWithBids: %+v", handsWithBids)
	// fmt.Println("result is", result)
}

func processFileContent(fileContent string) ([]HandWithBid, error) {
	var handsWithBids []HandWithBid
	lines := strings.Split(fileContent, "\n")
	for _, line := range lines {
		line = strings.Trim(line, " ")
		lineSplitted := strings.Split(line, " ")

		cards := lineSplitted[0]
		bid := lineSplitted[1]
		handWithBid, err := NewHandWithBid(cards, bid)
		if err != nil {
			return nil, err
		}
		handsWithBids = append(handsWithBids, *handWithBid)
	}

	return handsWithBids, nil
}
