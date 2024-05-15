package main

import (
	"fmt"
)

func isPalindrome(num int) bool {
	original := num
	reversed := 0

	// Reverse the number
	for num > 0 {
		digit := num % 10
		reversed = reversed*10 + digit
		num /= 10
	}

	// Check if the original number is equal to its reversed version
	return original == reversed
}

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&num)

	if isPalindrome(num) {
		fmt.Println(num, "is a palindrome")
	} else {
		fmt.Println(num, "is not a palindrome")
	}
}
