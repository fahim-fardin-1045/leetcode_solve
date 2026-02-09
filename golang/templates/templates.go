package main

import "fmt"

// Paste any helper structs or functions here
// Example: linked list
type ListNode struct {
    Val  int
    Next *ListNode
}

// Example helper function
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

// Problem-specific function
func yourFunction() {
    fmt.Println("Write your solution here")
}

func main() {
    yourFunction()
}
