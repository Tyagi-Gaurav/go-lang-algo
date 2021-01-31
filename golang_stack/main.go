package main

import (
	"fmt"
	"errors"
)

/*
Uses type assertions to implement a stack in GoLang. 
*/

type stack []interface{} 

func (s *stack) Push(v interface{}) {
	*s = append(*s, v)
}

/*
 Returning interface{} here is not a good strategy. Instead, we should 
 let the target type implement an interface that can accept this type as 
 parameter to extract the value.
*/
func (s *stack) Pop() (interface{}, error) {
	index := len(*s) - 1

	if index >= 0 {
		item := (*s)[index]
		*s = (*s)[0:index]
		return item, nil
	}

	return -1, errors.New("No more elements")
}

func main() {
	s := &stack{}

	s.Push(8)
	s.Push(20)

	item, _ := s.Pop()
	fmt.Printf("Popped Element from Stack: %d\n", item.(int))

	item, _ = s.Pop()
	fmt.Printf("Popped Element from Stack: %d\n", item.(int))

	_, err := s.Pop()
	if err != nil {
		fmt.Printf("No more elements on stack.")
	}
}
