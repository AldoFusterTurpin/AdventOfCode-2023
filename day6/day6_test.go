package main

import (
	"fmt"
	"testing"
)

func TestGetMultiplicationOfAllTheNumberOfWaysToBeatTheRecords(t *testing.T) {
	type testCase struct {
		times     []int
		distances []int
		expected  int
	}

	testCases := []testCase{
		{
			times:     []int{7, 15, 30},
			distances: []int{9, 40, 200},
			expected:  288,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			got := getMultiplicationOfAllTheNumberOfWaysToBeatTheRecords(tc.times, tc.distances)
			if got != tc.expected {
				t.Fatalf("expected %v, but got %v", tc.expected, got)
			}
		})
	}
}
