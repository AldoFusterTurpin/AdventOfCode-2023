package main

import (
	"testing"
)

func TestIsGamePossible(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		expected bool
	}

	testCases := []testCase{
		{
			name:     "Game 1 is possible",
			input:    "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			expected: true,
		},
		{
			name:     "Game 2 is possible",
			input:    "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			expected: true,
		},
		{
			name:     "Game 5 is possible",
			input:    "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			expected: true,
		},
		{
			name:     "Game 3 is not possible",
			input:    "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			expected: false,
		},
		{
			name:     "Game 4 is not possible",
			input:    "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := isGamePossible(tc.input)
			if got != tc.expected {
				t.Fatalf("expected %v, but got %v", tc.expected, got)
			}

		})
	}
}
