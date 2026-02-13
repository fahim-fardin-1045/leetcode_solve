package main

import "fmt"

func isValid(s string) bool {
	// Map closing brackets to their corresponding opening brackets
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	// Stack to store opening brackets
	stk := make([]rune, 0, len(s))

	for _, c := range s {
		// If the character is a closing bracket (exists as a key in our map)
		if open, found := pairs[c]; found {
			// Check if stack is empty OR if the top doesn't match the required opening bracket
			if len(stk) == 0 || stk[len(stk)-1] != open {
				return false
			}
			// Pop from the stack
			stk = stk[:len(stk)-1]
		} else {
			// It's an opening bracket, push it onto the stack
			stk = append(stk, c)
		}
	}

	// If the stack is empty, all brackets were matched correctly
	return len(stk) == 0
}

func main() {
	input := "{[()]}"
	fmt.Printf("Is '%s' valid? %v\n", input, isValid(input)) // Output: true
}
