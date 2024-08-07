package main

import (
	"strconv"
	"testing"
)

// This test is redundant after creating "Test_processFileContent" as NewHand() is never called
// in any other place outside of "processFileContent" (which is already covered by "Test_processFileContent"),
// but I prefer to keep it as it for now.
func TestNewHand(t *testing.T) {
	type testCase struct {
		cards        string
		wantHandType HandType
	}
	tests := []testCase{
		{
			cards:        "32T3K",
			wantHandType: OnePair,
		},
		{
			cards:        "T55J5",
			wantHandType: ThreeOfAkind,
		},
		{
			cards:        "KK677",
			wantHandType: TwoPair,
		},
		{
			cards:        "KTJJT",
			wantHandType: TwoPair,
		},
		{
			cards:        "QQQJA",
			wantHandType: ThreeOfAkind,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			hand := NewHand(tt.cards)
			handType := hand.handType

			if handType != tt.wantHandType {
				t.Errorf("expected %v, but got %v", tt.wantHandType, handType)
			}
		})
	}
}

func TestGetTotalWinnings(t *testing.T) {
	type testCase struct {
		hs   []HandWithBid
		want int
	}
	tests := []testCase{
		{
			hs: []HandWithBid{
				{
					hand: Hand{
						cards:    "32T3K",
						handType: OnePair,
					},
					bid: 765,
				},
				{
					hand: Hand{
						cards:    "KTJJT",
						handType: TwoPair,
					},
					bid: 220,
				},
				{
					hand: Hand{
						cards:    "KK677",
						handType: TwoPair,
					},
					bid: 28,
				},
				{
					hand: Hand{
						cards:    "T55J5",
						handType: ThreeOfAkind,
					},
					bid: 684,
				},
				{
					hand: Hand{
						cards:    "QQQJA",
						handType: ThreeOfAkind,
					},
					bid: 483,
				},
			},
			want: 6440,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := GetTotalWinnings(tt.hs)
			if got != tt.want {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}
}
