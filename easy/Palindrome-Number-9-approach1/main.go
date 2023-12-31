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

	return x == reverseNumber(x)
}

func reverseNumber(x int) int {
	reverse := 0
	for {
		reverse = reverse*10 + x%10
		x = x / 10
		if x == 0 {
			break
		}
	}

	return reverse
}
