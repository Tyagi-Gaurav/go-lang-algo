package main

import (
	"testing"
)

func TestBday(t *testing.T) {
	var tests = []struct {
		start int32
		end int32
		d int32
		e int32
	}{
		{20, 23, 6, 2},
		{13, 45, 3, 33},
	}
	for _, test := range tests {
		if a := Bday(test.start, test.end, test.d) ; a != test.e {
			t.Errorf("Test failed:  start: %d, end: %d, expected: %d, got: %d", test.start, test.end, test.e, a)
		}
	}
}
