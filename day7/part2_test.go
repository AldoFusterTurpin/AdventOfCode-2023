package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_getHandTypeWhenJokerIsPresent(t *testing.T) {
	type testCase struct {
		cards string
		want  HandType
	}

	tests := []testCase{
		{
			cards: "T55J5",
			want:  FourOfAkind,
		},
		{
			cards: "KTJJT",
			want:  FourOfAkind,
		},
		{
			cards: "QQQJA",
			want:  FourOfAkind,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			if got := getHandTypeWhenJokerIsPresent(tt.cards); got != tt.want {
				t.Errorf("getHandTypeWhenJokerIsPresent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getIndexes(t *testing.T) {
	type testCase struct {
		cards                 string
		targetChar            rune
		wantIndexesOfTarget   []int
		wantIndexesOfNoTarget []int
	}
	tests := []testCase{
		{
			cards:                 "32T3K",
			targetChar:            'J',
			wantIndexesOfTarget:   nil,
			wantIndexesOfNoTarget: []int{0, 1, 2, 3, 4},
		},
		{
			cards:                 "T55J5",
			targetChar:            'J',
			wantIndexesOfTarget:   []int{3},
			wantIndexesOfNoTarget: []int{0, 1, 2, 4},
		},
		{
			cards:                 "KK677",
			targetChar:            'J',
			wantIndexesOfTarget:   nil,
			wantIndexesOfNoTarget: []int{0, 1, 2, 3, 4},
		},
		{
			cards:                 "KTJJT",
			targetChar:            'J',
			wantIndexesOfTarget:   []int{2, 3},
			wantIndexesOfNoTarget: []int{0, 1, 4},
		},
		{
			cards:                 "QQQJA",
			targetChar:            'J',
			wantIndexesOfTarget:   []int{3},
			wantIndexesOfNoTarget: []int{0, 1, 2, 4},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			gotIndexesOfTarget, gotIndexesOfNoTarget := getIndexes(tt.cards, tt.targetChar)
			if !reflect.DeepEqual(gotIndexesOfTarget, tt.wantIndexesOfTarget) {
				t.Errorf("getIndexes() gotIndexesOfTarget = %v, want %v", gotIndexesOfTarget, tt.wantIndexesOfTarget)
			}
			if !reflect.DeepEqual(gotIndexesOfNoTarget, tt.wantIndexesOfNoTarget) {
				t.Errorf("getIndexes() gotIndexesOfNoTarget = %v, want %v", gotIndexesOfNoTarget, tt.wantIndexesOfNoTarget)
			}
		})
	}
}

func Test_getCharCounterWithoutJoker(t *testing.T) {
	type testCase struct {
		s    string
		want map[rune]int
	}
	tests := []testCase{
		{
			s:    "32T3K",
			want: map[rune]int{'3': 2, '2': 1, 'T': 1, 'K': 1},
		},
		{
			s:    "T55J5",
			want: map[rune]int{'T': 1, '5': 3},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			if got := getCharCounterWithoutJoker(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCharCounterWithoutJoker() = %v, want %v", got, tt.want)
			}
		})
	}
}
