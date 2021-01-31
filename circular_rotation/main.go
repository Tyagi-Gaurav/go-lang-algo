package main

import ("fmt"
	"strings"
)

func isCircularRotation(input string, testRotation string) bool {
	l1 := len(input)
	l2 := len(testRotation)

	if l1 == l2 {
		s1 := strings.Split(input, "")
		
		for i := 0; i < l1; i++ {
			permutation := strings.Join(append(s1[i:l1], s1[0:i]...), "")
			if testRotation == permutation {
				return true
			}
		}
	}

	return false
}

func main() {
	input := "ACTGACG"

	testRotation := "TGACGAC"
	fmt.Printf("Is %s circular rotation of %s: %t\n", testRotation, input,
		isCircularRotation(input, testRotation))

	testRotation = "TGACGAD"
	fmt.Printf("Is %s circular rotation of %s: %t\n", testRotation, input,
		isCircularRotation(input, testRotation))

	testRotation = "TGAD"
	fmt.Printf("Is %s circular rotation of %s: %t\n", testRotation, input,
		isCircularRotation(input, testRotation))

	input = "ABCD"
	testRotation = "CDAB"
	fmt.Printf("Is %s circular rotation of %s: %t\n", testRotation, input,
		isCircularRotation(input, testRotation))
}
