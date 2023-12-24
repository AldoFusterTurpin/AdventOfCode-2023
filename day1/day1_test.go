package main

import (
	"testing"
)

func TestGetCalibrationValue(t *testing.T) {
	type testCase struct {
		name        string
		input       string
		expected    string
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "expected 9",
			input:       "9vxfg",
			expected:    "99",
			expectedErr: nil,
		},
		{
			name:        "sample 38",
			input:       "pqr3stu8vwx",
			expected:    "38",
			expectedErr: nil,
		},
		{
			name:        "sample 77",
			input:       "treb7uchet",
			expected:    "77",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := getCalibrationValue(tc.input)

			if got != tc.expected {
				t.Fatalf("expected %v, but got %v", tc.expected, got)
			}

			if err != tc.expectedErr {
				t.Fatalf("expected error %v, but got error %v", tc.expectedErr, err)
			}
		})
	}
}

func TestReverseString(t *testing.T) {
	type testCase struct {
		name    string
		input   string
		want    string
		wantErr error
	}

	testCases := []testCase{
		{
			name:    "sample",
			input:   "9vxfg",
			want:    "gfxv9",
			wantErr: nil,
		},
		{
			name:    "sample 2",
			input:   "123",
			want:    "321",
			wantErr: nil,
		},
		{
			name:    "empty string",
			input:   "",
			want:    "",
			wantErr: nil,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := reverseString(tt.input)
			if err != tt.wantErr {
				t.Fatalf("expected error %v, but got %v", tt.wantErr, err)
			}
			if got != tt.want {
				t.Errorf("expected %v, but got %v", tt.want, got)
			}
		})
	}
}

func TestGetSumOfCalibrationValues(t *testing.T) {
	type testCase struct {
		name        string
		input       []string
		expected    int
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "expected 9",
			input:       []string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"},
			expected:    142,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := GetSumOfCalibrationValues(tc.input)

			if err != tc.expectedErr {
				t.Fatalf("expected error %v, but got error %v", tc.expectedErr, err)
			}

			if got != tc.expected {
				t.Fatalf("expected %v, but got %v", tc.expected, got)
			}
		})
	}
}

func TestGetSpelledDigitAtBeginningOfS(t *testing.T) {
	type testCase struct {
		name      string
		input     string
		expected  string
		expectedB bool
	}

	testCases := []testCase{
		{
			name:      "expected false as 9 is in digit form",
			input:     "9four",
			expected:  "",
			expectedB: false,
		},
		{
			name:      "expected false not enough letters",
			input:     "fiv",
			expected:  "",
			expectedB: false,
		},
		{
			name:      "expected 2",
			input:     "two1nine",
			expected:  "2",
			expectedB: true,
		},
		{
			name:      "expected 9",
			input:     "ninetwothree",
			expected:  "9",
			expectedB: true,
		},
		{
			name:      "expected 4",
			input:     "four247one",
			expected:  "4",
			expectedB: true,
		},
		{
			name:      "expected 1",
			input:     "one",
			expected:  "1",
			expectedB: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, b := getSpelledDigitAtBeginningOfS(tc.input, spelledNumbers())

			if b != tc.expectedB {
				t.Fatalf("expected b %v, but got %v", tc.expectedB, b)
			}

			if got != tc.expected {
				t.Fatalf("expected %v, but got %v", tc.expected, got)
			}
		})
	}
}

func TestFirstDigitFromStringPart2(t *testing.T) {
	type testCase struct {
		name                        string
		input                       string
		digitsSpelledOutWithLetters []string
		expected                    string
	}

	testCases := []testCase{
		{
			name:                        "expected 2",
			input:                       "two1nine",
			expected:                    "2",
			digitsSpelledOutWithLetters: spelledNumbers(),
		},
		{
			name:                        "expected 8",
			input:                       "eightwothree",
			expected:                    "8",
			digitsSpelledOutWithLetters: spelledNumbers(),
		},
		{
			name:                        "expected 1",
			input:                       "abcone2threexyz",
			expected:                    "1",
			digitsSpelledOutWithLetters: spelledNumbers(),
		},
		{
			name:                        "expected 2",
			input:                       "xtwone3four",
			expected:                    "2",
			digitsSpelledOutWithLetters: spelledNumbers(),
		},
		{
			name:                        "expected 4",
			input:                       "4nineeightseven2",
			expected:                    "4",
			digitsSpelledOutWithLetters: spelledNumbers(),
		},
		{
			name:                        "expected 4",
			input:                       "ruof3enowtx",
			expected:                    "4",
			digitsSpelledOutWithLetters: spelledNumbersReversed(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := firstDigitFromStringPart2(tc.input, tc.digitsSpelledOutWithLetters)
			if got != tc.expected {
				t.Fatalf("expected %v, but got %v", tc.expected, got)
			}
		})
	}
}

func TestGetCalibrationValuePart2(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		expected string
		wantErr  error
	}

	testCases := []testCase{
		{
			name:     "expected 29",
			input:    "two1nine",
			expected: "29",
		},
		{
			name:     "expected 8",
			input:    "eightwothree",
			expected: "83",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := getCalibrationValuePart2(tc.input)

			if err != tc.wantErr {
				t.Fatalf("expected error %v, but got %v", tc.wantErr, err)
			}

			if got != tc.expected {
				t.Fatalf("expected %v, but got %v", tc.expected, got)
			}
		})
	}
}
