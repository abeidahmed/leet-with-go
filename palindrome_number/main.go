package main

import (
	"fmt"
)

func splitInt(x int) []int {
	slice := []int{}

	for x > 0 {
		slice = append(slice, x%10)
		x = x / 10
	}

	return slice
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	splittedInt := splitInt(x)

	for i := 0; i < len(splittedInt); i++ {
		if splittedInt[i] != splittedInt[len(splittedInt)-i-1] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isPalindrome(1))
}
