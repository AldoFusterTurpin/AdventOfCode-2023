package main

import (
	"reflect"
	"strconv"
	"testing"
)

func Test_sortHandsWithBidsByStrength(t *testing.T) {
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
