package main

import (
	"fmt"
	"strings"
)

/*
Write a program that takes from standard input an expression without left parentheses and prints the equivalent infix expression with the 
parentheses inserted. For example, given the input:
1 + 2 ) * 3 - 4 ) * 5- 6 ) ) )
your program should print
( ( 1 + 2 ) * ( ( 3 -4 ) * ( 5 - 6 ) )
*/

func addBrackets(input string) string {
	stack := []string{}

	for i := 0 ; i < len(input); i++ {
		if input[i] != ' ' {
			if input[i] != ')' {
				stack = append(stack, string(input[i]))
			} else {
				l := len(stack)
				temp := stack[l-3:l]
				stack = append(stack[0:l-3], fmt.Sprintf("(%s)",strings.Join(temp, " ")))
			}
		}
	}

	return stack[0]
}

func main() {
	input := "1 + 2 ) * 3 - 4 ) * 5- 6 ) ) )"
	fmt.Printf("Without brackets: %s\nWith brackets: %s\n", input, addBrackets(input))

	input = "1 + 2) / 3 - 4)) * 5 - 6))"
	fmt.Printf("\nWithout brackets: %s\nWith brackets: %s\n", input, addBrackets(input))
}
