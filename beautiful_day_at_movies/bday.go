package main

import (
	"math"
)

func Bday(start int32, end int32, divisor int32) int32 {
	total := int32(0)
	for i:= start; i <= end; i++ {
		result := int64(math.Abs(float64(i - reverse(i)))) % int64(divisor)
		if result == 0 {
			total ++
		}
	}

	return total
}

func reverse(num int32) int32 {
	total:= int32(0)
	x := int32(0)

	for num != 0 {
		x = num % 10
		total = total * 10 + x
		num = num / 10
	}

	return total
}

func main() {
	
}
