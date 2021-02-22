package main

import (
	"math"
)

func Viral_load(n int32) int32 {
	c := int32(2)
	l := int32(2)
	s := int32(5)

	for i := int32(2); i <= n; i++ {
		s = l * 3
		l = int32(math.Floor(float64(s)/2))
		c = c + l
	}

	return c
}

func main() {
}
