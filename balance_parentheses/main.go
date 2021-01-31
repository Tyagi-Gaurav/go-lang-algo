package main

import (
	"fmt"
)

/*
Reads parenthesis from uses stack to determine if the set of parenthesis is balanced or not. Prints true if it is balanced and false otherwise.
*/

func verifyParenthesis(brack string) bool {
	c := 0
	
	for i := range(brack) {
		if brack[i] == '{' {
			c++
		} else {
			c--
		}
	}

	return c == 0
}


func main() {
	p1 := "{}{}{}{}"
	fmt.Printf("Are parenthesis '%s' balanced: %v\n", p1, verifyParenthesis(p1))

	p2 := "{}{}{}{"
	fmt.Printf("Are parenthesis '%s' balanced: %v\n", p2, verifyParenthesis(p2))

	p3 := "{{}{{{}}{}}}"
	fmt.Printf("Are parenthesis '%s' balanced: %v\n", p3, verifyParenthesis(p3))

	p4 := "{}}}"
	fmt.Printf("Are parenthesis '%s' balanced: %v\n", p4, verifyParenthesis(p4))
}
