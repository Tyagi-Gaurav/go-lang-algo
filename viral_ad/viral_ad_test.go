package main

import (
	"testing"
)

func TestViralLoad(t *testing.T) {
	var tests = []struct {
		input int32
		output int32
	}{
		{1, 2},
		{2, 5},
		{3, 9},
		{4, 15},
		{5, 24},
		{50, 2068789129},
	}

	for _, test := range tests {
		if o := Viral_load(test.input); test.output != o {
			t.Errorf("Test failed: input: %d, output: %d, expected: %d", test.input, o, test.output)
		}
	}
}
