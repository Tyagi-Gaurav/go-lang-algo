package main

func Circular_Rotation(a []int32, k int32, queries []int32) []int32 {
	l := int32(len(a))
	s := int32(0)

	for ; k > 0; k-- {
		s = s - 1
		if s < 0 {
			s = l - 1
		}
	}

	result := make([]int32, len(queries))

	for i, _ := range queries {
		index := (s + queries[i]) % l
		result[i] = a[index]
	}

	return result
}

func main() {
}
