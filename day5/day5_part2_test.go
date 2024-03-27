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
			input:                     sample,
			expected:                  46,
			desiredNumberOfRangePairs: -1,
		},
		{
			input:                     input2,
			expected:                  957727150,
			desiredNumberOfRangePairs: 1000,
		},
		{
			input:                     input,
			expected:                  1493866,
			desiredNumberOfRangePairs: 10_000,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			t.Parallel()
			got := GetLowestLocationOfSeedPairsConcurrent(tc.input, tc.desiredNumberOfRangePairs)
			if got != tc.expected {
				t.Fatalf("expected %v, but got %v", tc.expected, got)
			}
		})
	}
}
