package main

import "fmt"

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

	s := toSlice(x)
	lastIndex := len(s) - 1
	for i := 0; i < lastIndex; i++ {
		if s[i] != s[lastIndex] {
			return false
		}
		lastIndex--
	}

	return true
}

func toSlice(x int) []int {
	var s []int
	for {
		s = append(s, x%10)
		x = x / 10
		if x < 1 {
			break
		}
	}

	return s
}
