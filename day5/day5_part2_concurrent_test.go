package main

import (
	"fmt"
	"testing"
)

func TestGetLowestLocationOfSeedPairsConcurrent(t *testing.T) {
	type testCase struct {
		input                     string
		desiredNumberOfRangePairs int
		expected                  int
	}

	testCases := []testCase{
		{
			input:    sample,
			expected: 46,
			// we don't want to break small ranges into subranges.
			desiredNumberOfRangePairs: -1,
		},
		{
			input:                     input,
			expected:                  1493866,
			desiredNumberOfRangePairs: 10_000,
		},
		{
			input:                     input2,
			expected:                  957727150,
			desiredNumberOfRangePairs: 10_000,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			got := GetLowestLocationOfSeedPairsConcurrent(tc.input, tc.desiredNumberOfRangePairs)
			if got != tc.expected {
				t.Fatalf("expected %v, but got %v", tc.expected, got)
			}
		})
	}
}
