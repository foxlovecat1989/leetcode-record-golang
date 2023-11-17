package main

import (
	"fmt"
	"strconv"
)

func main() {
	result1 := isPalindrome(121)
	result2 := isPalindrome(-121)
	result3 := isPalindrome(10)
	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
}

func isPalindrome(x int) bool {
	// e.g., -10
	if x < 0 {
		return false
	}
	// e.g., 3
	if x/10 <= 0 {
		return true
	}

	s := strconv.Itoa(x)
	for i := 0; i < len(s)/2; i++ {
		lastPointer := len(s) - 1 - i
		if s[i] != s[lastPointer] {
			return false
		}
	}

	return true
}
