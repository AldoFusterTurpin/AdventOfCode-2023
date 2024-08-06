package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_convertStringToNodes(t *testing.T) {
	type testCase struct {
		inputString string
		want        []Node
	}
	tests := []testCase{
		{
			inputString: `AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`,
			want: []Node{
				{Id: "AAA", LeftElement: "BBB", RightElement: "CCC"},
				{Id: "BBB", LeftElement: "DDD", RightElement: "EEE"},
				{Id: "CCC", LeftElement: "ZZZ", RightElement: "GGG"},
				{Id: "DDD", LeftElement: "DDD", RightElement: "DDD"},
				{Id: "EEE", LeftElement: "EEE", RightElement: "EEE"},
				{Id: "GGG", LeftElement: "GGG", RightElement: "GGG"},
				{Id: "ZZZ", LeftElement: "ZZZ", RightElement: "ZZZ"},
			},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			if got := convertStringToNodes(tt.inputString); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertStringToNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
