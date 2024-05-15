package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	nextDistinct := 0
	for current := 1; current < len(nums); current++ {
		if nums[current] != nums[nextDistinct] {
			nextDistinct++
			nums[nextDistinct] = nums[current]
		}
	}

	return nextDistinct + 1
}

func main() {
	A := []int{1, 2, 2, 3, 4, 5, 5, 6, 7, 7, 8}
	length := removeDuplicates(A)
	fmt.Println("Length of the sorted array of distinct elements:", length)
	fmt.Println("Array after removing duplicates:", A[:length])
}

//Good code, but use a hashmap instead 