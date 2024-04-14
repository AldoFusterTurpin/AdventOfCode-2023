package main

import (
	"reflect"
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

func Test_processFileContent(t *testing.T) {
	type testCase struct {
		fileContent     string
		wantHandWithBid []HandWithBid
		wantErr         bool
	}

	tests := []testCase{
		{
			fileContent: `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`,
			wantHandWithBid: []HandWithBid{
				{
					hand: Hand{
						cards:    "32T3K",
						handType: OnePair,
					},
					bid: 765,
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
						cards:    "KK677",
						handType: TwoPair,
					},
					bid: 28,
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
						cards:    "QQQJA",
						handType: ThreeOfAkind,
					},
					bid: 483,
				},
			},
		},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got, err := getHandsWithBidsFromFileContent(tt.fileContent)
			if err != nil && !tt.wantErr {
				t.Fatalf("got error = %v, but wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && tt.wantErr {
				t.Fatalf("got not error, but wanted error %v", tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.wantHandWithBid) {
				t.Errorf("expected %v, but got %v", tt.wantHandWithBid, got)
			}
		})
	}
}

func Test_orderByStrength(t *testing.T) {
	type testCase struct {
		handWithBid     []HandWithBid
		wantHandWithBid []HandWithBid
	}

	tests := []testCase{
		{
			handWithBid: []HandWithBid{
				{
					hand: Hand{
						cards:    "32T3K",
						handType: OnePair,
					},
					bid: 765,
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
						cards:    "KK677",
						handType: TwoPair,
					},
					bid: 28,
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
						cards:    "QQQJA",
						handType: ThreeOfAkind,
					},
					bid: 483,
				},
			},
			wantHandWithBid: []HandWithBid{
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
		},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			sortHandsWithBidsByStrength(tt.handWithBid)
			if !reflect.DeepEqual(tt.wantHandWithBid, tt.handWithBid) {
				t.Errorf("expected %v, but got %v", tt.wantHandWithBid, tt.handWithBid)
			}
		})
	}
}
