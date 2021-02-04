package main

import (
	"testing"
)

func TestClosestPair(t *testing.T) {
	var tests = []struct {
		input []float64
		o1 float64
		o2 float64
	}{
		{[]float64{1.2, 1.0, 3.5, 5.8, 0.4, 10.8, 1.3, 6.4}, 1.2, 1.3},
		{[]float64{6.2, 8.34, 3.5, 15.8, 9.4, 10.8, 1.3, 6.4}, 6.2, 6.4},
		{[]float64{6.2, 6.4}, 6.2, 6.4},
	}

	for _, test := range tests {
		if a, b := ClosestPair(test.input); a != test.o1 && b != test.o2 {
			t.Errorf("Test failed: %v input, expected. (%f, %f), got: (%f, %f)", test.input, test.o1, test.o2, a, b)
		}
	}
}
