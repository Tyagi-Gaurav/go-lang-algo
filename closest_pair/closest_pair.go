package main

import (
	"sort"
	"math"
)

func ClosestPair(data []float64) (float64, float64) {
	sort.Float64s(data)
	a := data[0]
	b := data[1]
	min := math.Abs(a - b)

	for i:= 1; i < len(data) - 1; i++ {
		diff := math.Abs(data[i] - data[i+1])
		if diff < min {
			a = data[i]
			b = data[i+1]
			min = diff
		}
	}
	
	return a, b
}

func main() {
}
