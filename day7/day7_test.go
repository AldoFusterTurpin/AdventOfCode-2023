package main

import (
	"strconv"
	"testing"
)

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
