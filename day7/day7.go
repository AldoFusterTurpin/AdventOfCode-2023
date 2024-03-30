package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type HandType int

const (
	FiveOfAind HandType = iota
	FourOfAkind
	FullHouse
	ThreeOfAkind
	TwoPair
	OnePair
	HighCard
)

func (h HandType) String() string {
	switch h {
	case FiveOfAind:
		return "FiveOfAind"
	case FourOfAkind:
		return "FourOfAkind"
	case FullHouse:
		return "FullHouse"
	case ThreeOfAkind:
		return "ThreeOfAkind"
	case TwoPair:
		return "TwoPair"
	case OnePair:
		return "OnePair"
	case HighCard:
		return "HighCard"
	default:
		return "unknown"
	}
}

type Hand struct {
	cards    string
	handType HandType
}

func (h *Hand) setHandType() {
	for _, r := range h.cards {
		fmt.Print(string(r))
	}
	fmt.Println("\n__________________")
	h.handType = 0
}

func NewHand(cards string) *Hand {
	h := &Hand{cards: cards}
	h.setHandType()

	return h
}

type HandWithBid struct {
	hand Hand
	bid  int
}

func NewHandWithBid(cards, bid string) (*HandWithBid, error) {
	bidInt, err := strconv.Atoi(bid)
	if err != nil {
		return nil, err
	}

	handWithBid := &HandWithBid{
		hand: *NewHand(cards),
		bid:  bidInt,
	}

	return handWithBid, nil
}

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
	fmt.Println(handsWithBids)
	// fmt.Println("result is", result)
}

// 32T3K 765
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
