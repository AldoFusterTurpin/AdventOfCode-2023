package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
	}

	testCases := []testCase{
		{},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			got := fn()
			if got != tc.expected {
				t.Fatalf("expected %v, but got %v", tc.expected, got)
			}
		})
	}
}
