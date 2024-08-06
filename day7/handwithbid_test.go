package main

import (
	"reflect"
	"strconv"
	"testing"
)

func Test_getHandsWithBidsFromFileContent(t *testing.T) {
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
