package main

import (
	_ "embed"
	"testing"
)

//go:embed input_files/input.txt
var input string

//go:embed input_files/sample.txt
var sample string

func TestGetTotalScratchcardsPuntuation(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		expected int
	}

	testCases := []testCase{
		{
			name:     "sample",
			input:    sample,
			expected: 13,
		},
		{
			name:     "full input",
			input:    input,
			expected: 22897,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := getTotalScratchcardsPuntuation(tc.input)
			if got != tc.expected {
				t.Fatalf("expected %v, but got %v", tc.expected, got)
			}
		})
	}
}

func TestGetTotalScratchcards(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		expected int
	}

	testCases := []testCase{
		{
			name:     "sample",
			input:    sample,
			expected: 30,
		},
		{
			name:     "full input",
			input:    input,
			expected: 5095824,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := getFinalNumberOfScratchcards(tc.input)

			if got != tc.expected {
				t.Fatalf("expected %v, but got %v", tc.expected, got)
			}
		})
	}
}
