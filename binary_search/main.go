package main

import (
	"fmt"
)

func binary_search(arr []int, key int) (bool) {
	a := 0
	b := len(arr)

	for {
		if a > b || a >= len(arr) {
			return false
		}
		
		mid := (a + b)/2

		if arr[mid] == key {
			return true
		} else if arr[mid] > key {
			b = mid-1
		} else {
			a = mid+1
		}
	}
}

func main() {
	arr := []int{10, 24, 29, 35, 37, 41, 43, 45, 52, 54}

	key := 29
	fmt.Printf("Is key %d available in array: %v\n", key, binary_search(arr, key))

	key = 43
	fmt.Printf("Is key %d available in array: %v\n", key, binary_search(arr, key))

	key = 39
	fmt.Printf("Is key %d available in array: %v\n", key, binary_search(arr, key))

	key = 1
	fmt.Printf("Is key %d available in array: %v\n", key, binary_search(arr, key))

	key = 80
	fmt.Printf("Is key %d available in array: %v\n", key, binary_search(arr, key))

	key = -1
	fmt.Printf("Is key %d available in array: %v\n", key, binary_search(arr, key))
}
