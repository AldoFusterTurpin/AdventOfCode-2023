package main

import (
	"testing"

	_ "embed"
)

//go:embed input_files/sample.txt
var input string

func Test_getResult(t *testing.T) {
	type testCase struct {
		name string
		s    string
		want int
	}
	tests := []testCase{
		{
			name: "sample",
			s:    input,
			want: 5905,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getResult(tt.s); got != tt.want {
				t.Errorf("getResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
