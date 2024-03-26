package main

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed input_files/input.txt
var input string

//go:embed input_files/input2.txt
var input2 string

//go:embed input_files/sample.txt
var sample string

func TestGetDestinationValue(t *testing.T) {
	type testCase struct {
		source   int
		m        MapFromSourceToDestination
		expected int
	}

	testCases := []testCase{
		{
			m: MapFromSourceToDestination{
				from:        "seed",
				destination: "soil",
				ranges: []Range{
					{
						sourceRangeStart:      98,
						destinationRangeStart: 50,
						rangeLength:           2,
					},
				},
			},
			source:   98,
			expected: 50,
		},
		{
			m: MapFromSourceToDestination{
				from:        "seed",
				destination: "soil",
				ranges: []Range{
					{
						sourceRangeStart:      98,
						destinationRangeStart: 50,
						rangeLength:           2,
					},
				},
			},
			source:   99,
			expected: 51,
		},
		{
			m: MapFromSourceToDestination{
				from:        "seed",
				destination: "soil",
				ranges: []Range{
					{
						sourceRangeStart:      98,
						destinationRangeStart: 50,
						rangeLength:           2,
					},
				},
			},
			source:   100,
			expected: 100,
		},
		{
			m: MapFromSourceToDestination{
				from:        "seed",
				destination: "soil",
				ranges: []Range{
					{
						sourceRangeStart:      98,
						destinationRangeStart: 50,
						rangeLength:           2,
					},
				},
			},
			source:   97,
			expected: 97,
		},
		{
			m: MapFromSourceToDestination{
				from:        "seed",
				destination: "soil",
				ranges: []Range{
					{
						sourceRangeStart:      50,
						destinationRangeStart: 52,
						rangeLength:           48,
					},
				},
			},
			source:   50,
			expected: 52,
		},
		{
			m: MapFromSourceToDestination{
				from:        "seed",
				destination: "soil",
				ranges: []Range{
					{
						sourceRangeStart:      50,
						destinationRangeStart: 52,
						rangeLength:           48,
					},
				},
			},
			source:   51,
			expected: 53,
		},
		{
			m: MapFromSourceToDestination{
				from:        "seed",
				destination: "soil",
				ranges: []Range{
					{
						sourceRangeStart:      50,
						destinationRangeStart: 52,
						rangeLength:           48,
					},
				},
			},
			source:   52,
			expected: 54,
		},
		{
			m: MapFromSourceToDestination{
				from:        "seed",
				destination: "soil",
				ranges: []Range{
					{
						sourceRangeStart:      50,
						destinationRangeStart: 52,
						rangeLength:           48,
					},
				},
			},
			source:   53,
			expected: 55,
		},
		{
			m: MapFromSourceToDestination{
				from:        "seed",
				destination: "soil",
				ranges: []Range{
					{
						sourceRangeStart:      50,
						destinationRangeStart: 52,
						rangeLength:           48,
					},
				},
			},
			source:   49,
			expected: 49,
		},
		{
			m: MapFromSourceToDestination{
				from:        "seed",
				destination: "soil",
				ranges: []Range{
					{
						sourceRangeStart:      50,
						destinationRangeStart: 52,
						rangeLength:           48,
					},
				},
			},
			source:   97,
			expected: 99,
		},
		{
			m: MapFromSourceToDestination{
				from:        "seed",
				destination: "soil",
				ranges: []Range{
					{
						sourceRangeStart:      50,
						destinationRangeStart: 52,
						rangeLength:           48,
					},
				},
			},
			source:   98,
			expected: 98,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			got := tc.m.GetDestinationValue(tc.source)
			if got != tc.expected {
				t.Fatalf("expected %v, but got %v", tc.expected, got)
			}
		})
	}
}

func TestAlmanac_GetLowestLocationOfAllSeeds(t *testing.T) {
	type testCase struct {
		almanac  Almanac
		expected int
	}

	testCases := []testCase{
		{
			almanac: Almanac{
				seedsToBePlanted: []int{79, 14, 55, 13},
				maps: []MapFromSourceToDestination{
					{
						from:        "seed",
						destination: "soil",
						ranges: []Range{
							{
								destinationRangeStart: 50,
								sourceRangeStart:      98,
								rangeLength:           2,
							},
							{
								destinationRangeStart: 52,
								sourceRangeStart:      50,
								rangeLength:           48,
							},
						},
					},
					{
						from:        "soil",
						destination: "fertilizer",
						ranges: []Range{
							{
								destinationRangeStart: 0,
								sourceRangeStart:      15,
								rangeLength:           37,
							},
							{
								destinationRangeStart: 37,
								sourceRangeStart:      52,
								rangeLength:           2,
							},
							{
								destinationRangeStart: 39,
								sourceRangeStart:      0,
								rangeLength:           15,
							},
						},
					},
					{
						from:        "fertilizer",
						destination: "water",
						ranges: []Range{
							{
								destinationRangeStart: 49,
								sourceRangeStart:      53,
								rangeLength:           8,
							},
							{
								destinationRangeStart: 0,
								sourceRangeStart:      11,
								rangeLength:           42,
							},
							{
								destinationRangeStart: 42,
								sourceRangeStart:      0,
								rangeLength:           7,
							},
							{
								destinationRangeStart: 57,
								sourceRangeStart:      7,
								rangeLength:           4,
							},
						},
					},
					{
						from:        "water",
						destination: "light",
						ranges: []Range{
							{
								destinationRangeStart: 88,
								sourceRangeStart:      18,
								rangeLength:           7,
							},
							{
								destinationRangeStart: 18,
								sourceRangeStart:      25,
								rangeLength:           70,
							},
						},
					},
					{
						from:        "light",
						destination: "temperature",
						ranges: []Range{
							{
								destinationRangeStart: 45,
								sourceRangeStart:      77,
								rangeLength:           23,
							},
							{
								destinationRangeStart: 81,
								sourceRangeStart:      45,
								rangeLength:           19,
							},
							{
								destinationRangeStart: 68,
								sourceRangeStart:      64,
								rangeLength:           13,
							},
						},
					},
					{
						from:        "temperature",
						destination: "humidity",
						ranges: []Range{
							{
								destinationRangeStart: 0,
								sourceRangeStart:      69,
								rangeLength:           1,
							},
							{
								destinationRangeStart: 1,
								sourceRangeStart:      0,
								rangeLength:           69,
							},
						},
					},
					{
						from:        "humidity",
						destination: "location",
						ranges: []Range{
							{
								destinationRangeStart: 60,
								sourceRangeStart:      56,
								rangeLength:           37,
							},
							{
								destinationRangeStart: 56,
								sourceRangeStart:      93,
								rangeLength:           4,
							},
						},
					},
				},
			},
			expected: 35,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			got := tc.almanac.getLowestLocationOfAllSeeds()
			if got != tc.expected {
				t.Fatalf("expected %v, but got %v", tc.expected, got)
			}
		})
	}
}

func TestGetLowestLocationOfAllSeeds(t *testing.T) {
	type testCase struct {
		input    string
		expected int
	}

	testCases := []testCase{
		{
			input:    sample,
			expected: 35,
		},
		{
			input:    input,
			expected: 174137457,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			got := GetLowestLocationOfAllSeeds(tc.input)
			if got != tc.expected {
				t.Fatalf("expected %v, but got %v", tc.expected, got)
			}
		})
	}
}
