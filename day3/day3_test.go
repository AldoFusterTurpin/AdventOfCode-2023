package main

import (
	_ "embed"
	"reflect"
	"testing"
)

//go:embed input_files/input.txt
var input string

func TestGetSumOfPartNumbers(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		expected int
	}

	testCases := []testCase{
		{
			name: "sample engine schematic",
			input: `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`,
			expected: 4361,
		},
		{
			name:     "full input engine schematic",
			input:    input,
			expected: 535078,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := getSumOfPartNumbers(tc.input)
			if got != tc.expected {
				t.Fatalf("expected %v, but got %v", tc.expected, got)
			}
		})
	}
}

func TestGetMapOfPartNumbers(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		expected map[Coordinate]int
	}

	testCases := []testCase{
		{
			name: "sample engine schematic",
			input: `467..114..
		...*......
		..35..633.
		......#...
		617*......
		.....+.58.
		..592.....
		......755.
		...$.*....
		.664.598..`,
			expected: map[Coordinate]int{
				{row: 0, col: 0}: 467,
				{row: 0, col: 1}: 467,
				{row: 0, col: 2}: 467,
				{row: 2, col: 2}: 35,
				{row: 2, col: 3}: 35,
				{row: 2, col: 6}: 633,
				{row: 2, col: 7}: 633,
				{row: 2, col: 8}: 633,
				{row: 4, col: 0}: 617,
				{row: 4, col: 1}: 617,
				{row: 4, col: 2}: 617,
				{row: 6, col: 2}: 592,
				{row: 6, col: 3}: 592,
				{row: 6, col: 4}: 592,
				{row: 7, col: 6}: 755,
				{row: 7, col: 7}: 755,
				{row: 7, col: 8}: 755,
				{row: 9, col: 1}: 664,
				{row: 9, col: 2}: 664,
				{row: 9, col: 3}: 664,
				{row: 9, col: 5}: 598,
				{row: 9, col: 6}: 598,
				{row: 9, col: 7}: 598,
			},
		},
		{
			name: "sample engine schematic 2",
			input: `..529
...*.
.380.
.....
.....`,
			expected: map[Coordinate]int{
				{row: 0, col: 2}: 529,
				{row: 0, col: 3}: 529,
				{row: 0, col: 4}: 529,
				{row: 2, col: 1}: 380,
				{row: 2, col: 2}: 380,
				{row: 2, col: 3}: 380,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := getMapOfPartNumbers(tc.input)

			areEqual := reflect.DeepEqual(got, tc.expected)
			if !areEqual {
				t.Fatalf("expected\n%v,\nbut got\n%v", tc.expected, got)
			}
		})
	}
}

func TestGetSumOfGearRatios(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		expected int
	}

	testCases := []testCase{
		{
			name: "sample engine schematic",
			input: `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`,
			expected: 467835,
		},
		{
			name:     "full input engine schematic",
			input:    input,
			expected: 75312571,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := getSumOfGearRatios(tc.input)
			if got != tc.expected {
				t.Fatalf("expected %v, but got %v", tc.expected, got)
			}
		})
	}
}
